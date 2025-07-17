// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEffectiveOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListEffectiveOrdersRequest
	GetTid() *int64
}

type ListEffectiveOrdersRequest struct {
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListEffectiveOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEffectiveOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListEffectiveOrdersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListEffectiveOrdersRequest) SetTid(v int64) *ListEffectiveOrdersRequest {
	s.Tid = &v
	return s
}

func (s *ListEffectiveOrdersRequest) Validate() error {
	return dara.Validate(s)
}
