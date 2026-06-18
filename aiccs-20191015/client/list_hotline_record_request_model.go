// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *ListHotlineRecordRequest
	GetCallId() *string
	SetClientToken(v string) *ListHotlineRecordRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListHotlineRecordRequest
	GetInstanceId() *string
}

type ListHotlineRecordRequest struct {
	// Session ID. This corresponds to the acid in WebSocket after an inbound call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100365558
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Unique customer request ID. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// AICCS instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListHotlineRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordRequest) GetCallId() *string {
	return s.CallId
}

func (s *ListHotlineRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListHotlineRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHotlineRecordRequest) SetCallId(v string) *ListHotlineRecordRequest {
	s.CallId = &v
	return s
}

func (s *ListHotlineRecordRequest) SetClientToken(v string) *ListHotlineRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordRequest) SetInstanceId(v string) *ListHotlineRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHotlineRecordRequest) Validate() error {
	return dara.Validate(s)
}
