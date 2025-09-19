// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCancelIgnoreSuspEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemark(v string) *OperationCancelIgnoreSuspEventRequest
	GetRemark() *string
	SetSecurityEventIds(v []*int64) *OperationCancelIgnoreSuspEventRequest
	GetSecurityEventIds() []*int64
}

type OperationCancelIgnoreSuspEventRequest struct {
	// The remarks.
	//
	// example:
	//
	// remark text
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IDs of alert events.
	//
	// This parameter is required.
	SecurityEventIds []*int64 `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s OperationCancelIgnoreSuspEventRequest) String() string {
	return dara.Prettify(s)
}

func (s OperationCancelIgnoreSuspEventRequest) GoString() string {
	return s.String()
}

func (s *OperationCancelIgnoreSuspEventRequest) GetRemark() *string {
	return s.Remark
}

func (s *OperationCancelIgnoreSuspEventRequest) GetSecurityEventIds() []*int64 {
	return s.SecurityEventIds
}

func (s *OperationCancelIgnoreSuspEventRequest) SetRemark(v string) *OperationCancelIgnoreSuspEventRequest {
	s.Remark = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventRequest) SetSecurityEventIds(v []*int64) *OperationCancelIgnoreSuspEventRequest {
	s.SecurityEventIds = v
	return s
}

func (s *OperationCancelIgnoreSuspEventRequest) Validate() error {
	return dara.Validate(s)
}
