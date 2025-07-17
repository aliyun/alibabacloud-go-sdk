// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogic(v bool) *ListColumnsRequest
	GetLogic() *bool
	SetTableId(v string) *ListColumnsRequest
	GetTableId() *string
	SetTid(v int64) *ListColumnsRequest
	GetTid() *int64
}

type ListColumnsRequest struct {
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The ID of the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the table ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39281****
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListColumnsRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListColumnsRequest) GetTableId() *string {
	return s.TableId
}

func (s *ListColumnsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListColumnsRequest) SetLogic(v bool) *ListColumnsRequest {
	s.Logic = &v
	return s
}

func (s *ListColumnsRequest) SetTableId(v string) *ListColumnsRequest {
	s.TableId = &v
	return s
}

func (s *ListColumnsRequest) SetTid(v int64) *ListColumnsRequest {
	s.Tid = &v
	return s
}

func (s *ListColumnsRequest) Validate() error {
	return dara.Validate(s)
}
