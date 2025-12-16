// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTableFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *ListDataSourceTableFieldsRequest
	GetParams() *string
	SetRawType(v bool) *ListDataSourceTableFieldsRequest
	GetRawType() *bool
}

type ListDataSourceTableFieldsRequest struct {
	// The parameters of the data source. The value of the params parameter is a JSON string. The value must be URL-encoded.
	//
	// Different types of data sources use different parameters. For more information, see the following sections of the "DataSource" topic:
	//
	// 	- [rds](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [polardb](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [odps](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [mysql](https://help.aliyun.com/document_detail/173627.html)
	//
	// 	- [drds](https://help.aliyun.com/document_detail/173627.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// Specifies whether to return the original field types of the data source.
	//
	// example:
	//
	// false
	RawType *bool `json:"rawType,omitempty" xml:"rawType,omitempty"`
}

func (s ListDataSourceTableFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTableFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsRequest) GetParams() *string {
	return s.Params
}

func (s *ListDataSourceTableFieldsRequest) GetRawType() *bool {
	return s.RawType
}

func (s *ListDataSourceTableFieldsRequest) SetParams(v string) *ListDataSourceTableFieldsRequest {
	s.Params = &v
	return s
}

func (s *ListDataSourceTableFieldsRequest) SetRawType(v bool) *ListDataSourceTableFieldsRequest {
	s.RawType = &v
	return s
}

func (s *ListDataSourceTableFieldsRequest) Validate() error {
	return dara.Validate(s)
}
