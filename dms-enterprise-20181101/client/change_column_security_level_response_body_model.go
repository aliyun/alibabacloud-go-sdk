// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecurityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ChangeColumnSecurityLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChangeColumnSecurityLevelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ChangeColumnSecurityLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeColumnSecurityLevelResponseBody
	GetSuccess() *bool
}

type ChangeColumnSecurityLevelResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E103C5F9-DE47-53F2-BF34-D71DF38F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeColumnSecurityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecurityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecurityLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeColumnSecurityLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChangeColumnSecurityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeColumnSecurityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeColumnSecurityLevelResponseBody) SetErrorCode(v string) *ChangeColumnSecurityLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeColumnSecurityLevelResponseBody) SetErrorMessage(v string) *ChangeColumnSecurityLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChangeColumnSecurityLevelResponseBody) SetRequestId(v string) *ChangeColumnSecurityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeColumnSecurityLevelResponseBody) SetSuccess(v bool) *ChangeColumnSecurityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeColumnSecurityLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
