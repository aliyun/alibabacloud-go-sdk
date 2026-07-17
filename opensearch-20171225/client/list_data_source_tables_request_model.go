// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *ListDataSourceTablesRequest
	GetParams() *string
}

type ListDataSourceTablesRequest struct {
	// The parameters for the data source. The value must be a URL-encoded JSON string.
	//
	// The parameters vary based on the data source type. For more information, see:
	//
	// - [rds](https://help.aliyun.com/document_detail/170005.html)
	//
	// - [polardb](https://help.aliyun.com/document_detail/170005.html)
	//
	// - [odps](https://help.aliyun.com/document_detail/170005.html)
	//
	// - [mysql](https://help.aliyun.com/document_detail/173627.html)
	//
	// - [drds](https://help.aliyun.com/document_detail/173627.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// -
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ListDataSourceTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesRequest) GetParams() *string {
	return s.Params
}

func (s *ListDataSourceTablesRequest) SetParams(v string) *ListDataSourceTablesRequest {
	s.Params = &v
	return s
}

func (s *ListDataSourceTablesRequest) Validate() error {
	return dara.Validate(s)
}
