// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallInfos(v string) *InsertTaskDetailRequest
	GetCallInfos() *string
	SetInstanceId(v string) *InsertTaskDetailRequest
	GetInstanceId() *string
	SetOutboundTaskId(v int64) *InsertTaskDetailRequest
	GetOutboundTaskId() *int64
}

type InsertTaskDetailRequest struct {
	// This parameter is required.
	CallInfos *string `json:"CallInfos,omitempty" xml:"CallInfos,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	OutboundTaskId *int64 `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
}

func (s InsertTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailRequest) GetCallInfos() *string {
	return s.CallInfos
}

func (s *InsertTaskDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InsertTaskDetailRequest) GetOutboundTaskId() *int64 {
	return s.OutboundTaskId
}

func (s *InsertTaskDetailRequest) SetCallInfos(v string) *InsertTaskDetailRequest {
	s.CallInfos = &v
	return s
}

func (s *InsertTaskDetailRequest) SetInstanceId(v string) *InsertTaskDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *InsertTaskDetailRequest) SetOutboundTaskId(v int64) *InsertTaskDetailRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *InsertTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
