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
	SetProductCode(v string) *GetCallDetailRecordRequest
	GetProductCode() *string
	SetSessionId(v string) *GetCallDetailRecordRequest
	GetSessionId() *string
}

type GetCallDetailRecordRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The product code.
	//
	// example:
	//
	// OUTBOUND_BOT
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The call session ID.
	//
	// example:
	//
	// job-0b84bf6f-73dc-4462-bd8f-916e3a34c419
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

func (s *GetCallDetailRecordRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCallDetailRecordRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCallDetailRecordRequest) SetInstanceId(v string) *GetCallDetailRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCallDetailRecordRequest) SetProductCode(v string) *GetCallDetailRecordRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCallDetailRecordRequest) SetSessionId(v string) *GetCallDetailRecordRequest {
	s.SessionId = &v
	return s
}

func (s *GetCallDetailRecordRequest) Validate() error {
	return dara.Validate(s)
}
