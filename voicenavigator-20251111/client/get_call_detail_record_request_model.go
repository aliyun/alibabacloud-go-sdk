// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDetailRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetCallDetailRecordRequest
	GetInstanceId() *string
	SetSessionId(v string) *GetCallDetailRecordRequest
	GetSessionId() *string
}

type GetCallDetailRecordRequest struct {
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9fe19f64-a290-4a73-a3ed-5f7d21d44ecb
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetCallDetailRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordRequest) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCallDetailRecordRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCallDetailRecordRequest) SetInstanceId(v string) *GetCallDetailRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCallDetailRecordRequest) SetSessionId(v string) *GetCallDetailRecordRequest {
	s.SessionId = &v
	return s
}

func (s *GetCallDetailRecordRequest) Validate() error {
	return dara.Validate(s)
}
