// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenariosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListScenariosRequest
	GetTid() *int64
}

type ListScenariosRequest struct {
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListScenariosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScenariosRequest) GoString() string {
	return s.String()
}

func (s *ListScenariosRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListScenariosRequest) SetTid(v int64) *ListScenariosRequest {
	s.Tid = &v
	return s
}

func (s *ListScenariosRequest) Validate() error {
	return dara.Validate(s)
}
