// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceCustomUserIdByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceCustomUserIdByUserIdResponseBody
	GetCode() *string
	SetData(v string) *QueryFaceCustomUserIdByUserIdResponseBody
	GetData() *string
	SetErrorMessage(v string) *QueryFaceCustomUserIdByUserIdResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceCustomUserIdByUserIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceCustomUserIdByUserIdResponseBody
	GetSuccess() *bool
}

type QueryFaceCustomUserIdByUserIdResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The user ID in the caller\\"s system.
	//
	// example:
	//
	// TestCustomUserId
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetCode(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetData(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetErrorMessage(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetRequestId(v string) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) SetSuccess(v bool) *QueryFaceCustomUserIdByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
