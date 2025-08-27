// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorSyncPayStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CooperatorSyncPayStatusResponseBody
	GetCode() *string
	SetMessage(v string) *CooperatorSyncPayStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *CooperatorSyncPayStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CooperatorSyncPayStatusResponseBody
	GetSuccess() *bool
}

type CooperatorSyncPayStatusResponseBody struct {
	// example:
	//
	// System.Error
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CooperatorSyncPayStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CooperatorSyncPayStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CooperatorSyncPayStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *CooperatorSyncPayStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CooperatorSyncPayStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CooperatorSyncPayStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CooperatorSyncPayStatusResponseBody) SetCode(v string) *CooperatorSyncPayStatusResponseBody {
	s.Code = &v
	return s
}

func (s *CooperatorSyncPayStatusResponseBody) SetMessage(v string) *CooperatorSyncPayStatusResponseBody {
	s.Message = &v
	return s
}

func (s *CooperatorSyncPayStatusResponseBody) SetRequestId(v string) *CooperatorSyncPayStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CooperatorSyncPayStatusResponseBody) SetSuccess(v bool) *CooperatorSyncPayStatusResponseBody {
	s.Success = &v
	return s
}

func (s *CooperatorSyncPayStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
