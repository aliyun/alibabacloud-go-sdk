// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateHostGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateHostGroupResponseBody
	GetErrorMessage() *string
	SetHostGroupId(v int64) *CreateHostGroupResponseBody
	GetHostGroupId() *int64
	SetRequestId(v string) *CreateHostGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHostGroupResponseBody
	GetSuccess() *bool
}

type CreateHostGroupResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 123
	HostGroupId *int64 `json:"hostGroupId,omitempty" xml:"hostGroupId,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateHostGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateHostGroupResponseBody) GetHostGroupId() *int64 {
	return s.HostGroupId
}

func (s *CreateHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHostGroupResponseBody) SetErrorCode(v string) *CreateHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetErrorMessage(v string) *CreateHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetHostGroupId(v int64) *CreateHostGroupResponseBody {
	s.HostGroupId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetRequestId(v string) *CreateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetSuccess(v bool) *CreateHostGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHostGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
