// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogic(v bool) *ListIndexesRequest
	GetLogic() *bool
	SetTableId(v string) *ListIndexesRequest
	GetTableId() *string
	SetTid(v int64) *ListIndexesRequest
	GetTid() *int64
}

type ListIndexesRequest struct {
	// Specifies whether the table is a logical table.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The ID of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 0
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListIndexesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexesRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListIndexesRequest) GetTableId() *string {
	return s.TableId
}

func (s *ListIndexesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListIndexesRequest) SetLogic(v bool) *ListIndexesRequest {
	s.Logic = &v
	return s
}

func (s *ListIndexesRequest) SetTableId(v string) *ListIndexesRequest {
	s.TableId = &v
	return s
}

func (s *ListIndexesRequest) SetTid(v int64) *ListIndexesRequest {
	s.Tid = &v
	return s
}

func (s *ListIndexesRequest) Validate() error {
	return dara.Validate(s)
}
