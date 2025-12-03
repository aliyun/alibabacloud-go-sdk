// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTestResultResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateTestResultResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *UpdateTestResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTestResultResponseBody
	GetSuccess() *bool
}

type UpdateTestResultResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTestResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTestResultResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTestResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTestResultResponseBody) SetErrorCode(v string) *UpdateTestResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTestResultResponseBody) SetErrorMsg(v string) *UpdateTestResultResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTestResultResponseBody) SetRequestId(v string) *UpdateTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTestResultResponseBody) SetSuccess(v bool) *UpdateTestResultResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTestResultResponseBody) Validate() error {
	return dara.Validate(s)
}
