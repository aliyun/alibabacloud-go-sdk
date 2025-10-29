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
	// The catalog entity ID. Currently supports dlf and starrocks types. You can refer to the results returned by the ListCatalogs operation and the [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// 	- For the DLF type, the format is `dlf-catalog::catalog_id`.
	//
	// 	- For the StarRocks type, the format is `starrocks-catalog:(instance_id|encoded_jdbc_url):catalog_name`.
	//
	// > \\
	//
	// `catalog_id`: The ID of the DLF catalog.\\
	//
	// `instance_id`: The instance ID, required if the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string. Required if the data source is registered in connection string mode.\\
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
