package main // entry postring of program
// indentifies an executable program

import (
	"fmt" //built-in package provided by Go that provides input/output utility functions.
)

const MAX_TENANTS = 14

var tenants map[string]string
var tenants_paid map[string]string 
var tenant_not_paid map[string]string

func main() {

    //declare size and capacity of tenants_paid and tenants

    tenants_paid = make(map[string]string, 14)
    tenants = make(map[string]string, 14)
    tenant_not_paid = make(map[string]string, 14)

    //initialize tenants
    
    //copies tenants to tenant_not_paid
    addTenant(&tenants, "Munene", "0712345678")
    
    hasPaid("Munene", "0741405735")
    
    fmt.Println(isPresent(&tenants_paid,"0741405735"))

}

//methods 

//copies tenants to tenant_not_paid
func copyCollections(tenant map[string]string) {

    for key, value := range tenant {
        tenant_not_paid[key] = value    
    }
}

// checks if present in certain collection
func isPresent(collection *map[string]string, phone_number string) bool {
	var found bool
	found = false

	for key := range *collection {
		if key == phone_number {
			found = true
			break
		}
	}
	return found
}

// adds a tenant to tenant_paid slice
func hasPaid(name string, phone_number string) {

    //if present in tenants 
	if isPresent(&tenants, phone_number) {
        //display
		fmt.Println(name,"with",phone_number,"has paid rent")
        //delete from the tenant_not_paid map
        delete(tenant_not_paid, phone_number)
        fmt.Println("Operation successful")
        addTenant(&tenants_paid, name, phone_number)
        fmt.Println("Added to tenants_paid map")

	}else{
        fmt.Println(phone_number,"not in use by any tenant")
    }
}

// shows all names in a certain collection
func showAll(collection *map[string]string) {
    for key, value := range *collection {
        fmt.Println(key, value)
    }
}

//add to a collection
func addTenant(collection *map[string]string, tenant_name string, phone_num string){

    //de-reference pointer
    map_pointer := *collection

    //if already present
    if isPresent(&map_pointer, phone_num){
        fmt.Println(tenant_name,"with",phone_num,"already present")
    }

    map_pointer[phone_num] = tenant_name
    fmt.Println(tenant_name,"with",phone_num,"added successfully")
    copyCollections(tenants)
}

