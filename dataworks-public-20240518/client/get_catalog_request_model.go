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
	// The ID of the data catalog entity. Currently, DLF and StarRocks types are supported. For more information, see the response of the ListCatalogs operation and [Concepts of metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	//
	// - For the DLF type, the format is `dlf-catalog::catalog_id`.
	//
	// - For the StarRocks type, the format is `starrocks-catalog:(instance_id|encoded_jdbc_url):catalog_name`.
	//
	// > Where:
	//
	// `catalog_id`: the DLF catalog ID.
	//
	// `instance_id`: the instance ID, which is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: the URL-encoded JDBC connection string, which is required when the data source is registered by using a connection string.
	//
	// `catalog_name`: the StarRocks catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf-catalog::catalog_id
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
