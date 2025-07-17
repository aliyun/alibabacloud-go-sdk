// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetCatalogRequest
	GetId() *string
}

type GetCatalogRequest struct {
	// Data catalog entity ID. Currently, only DLF and StarRocks types are supported. You can refer to the response of the ListCatalogs operation and [the description of metadata entity concepts.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// 	- For the DLF type, the format is `dlf-catalog::catalog_id`.
	//
	// 	- For the StarRocks type, the format is `starrocks-catalog:(instance_id|encoded_jdbc_url):catalog_name`.
	//
	// >  Parameter descriptions:\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `instance_id`: The instance ID, required for the data source registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The JDBC connection string that has been URL encoded, required for the data source registered via a connection string.\\
	//
	// `catalog_name`: The name of the StarRocks catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf-catalog:123456XXX:test_catalog
	//
	// starrocks-catalog:c-abc123xxx:default_catalog
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogRequest) GetId() *string {
	return s.Id
}

func (s *GetCatalogRequest) SetId(v string) *GetCatalogRequest {
	s.Id = &v
	return s
}

func (s *GetCatalogRequest) Validate() error {
	return dara.Validate(s)
}
