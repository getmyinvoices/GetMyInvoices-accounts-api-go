![get-my-invoices](https://www.getmyinvoices.com/wp-content/uploads/2016/04/logo_login.png)

# gogmi is a golang-library for getmyinvoices.com

## Already implemented

### V1
Here you can find the API-documentation: https://api.getmyinvoices.com/accounts/v1/doc/#
✔ list supplieres  
✔ get specific supplier  
✔ list invoices  
✘ upload new invoice  
✘ update invoice  
✔ get country list  
✘ add custom supplier  
✘ update custom supplier  
✘ delete custom supplier  
✘ get attachment list  
✘ upload one attachment
✘ delete one attachment  

### V2
Here you can find the API-documentation: https://api.getmyinvoices.com/accounts/v2/doc/#
✘ get company list  
✘ get one company  
✘ get document list  
✘ get one document  
✘ upload new document  
✘ update document  
✔ get country list  
✘ get custom company  
✘ add custom company  
✘ delete custom company  
✘ get attachment list  
✘ upload one attachment  
✘ delete one attachment  

## Getting started

```golang
client := gogmi.GMI{
    APIVersion: "v2",
    APIKey:     "your-API-Key",
}
suppliers, err := client.ListSuppliers()
if err != nil {
    t.Error(err)
}

// do something with suppliers
```

