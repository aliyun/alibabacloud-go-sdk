// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLhDagOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ChangeLhDagOwnerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChangeLhDagOwnerResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ChangeLhDagOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeLhDagOwnerResponseBody
	GetSuccess() *bool
}

type ChangeLhDagOwnerResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9997630E-1993-5E6D-9DF1-4EFEE755FE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeLhDagOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeLhDagOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeLhDagOwnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeLhDagOwnerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChangeLhDagOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeLhDagOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeLhDagOwnerResponseBody) SetErrorCode(v string) *ChangeLhDagOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeLhDagOwnerResponseBody) SetErrorMessage(v string) *ChangeLhDagOwnerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChangeLhDagOwnerResponseBody) SetRequestId(v string) *ChangeLhDagOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeLhDagOwnerResponseBody) SetSuccess(v bool) *ChangeLhDagOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeLhDagOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
