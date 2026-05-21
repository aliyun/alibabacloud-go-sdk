// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindLeaderInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UnBindLeaderInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *UnBindLeaderInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UnBindLeaderInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UnBindLeaderInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UnBindLeaderInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnBindLeaderInstanceResponseBody
	GetSuccess() *bool
}

type UnBindLeaderInstanceResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 288D934D-3F25-5BA6-9B50-7ABB3017A134
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindLeaderInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnBindLeaderInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindLeaderInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UnBindLeaderInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UnBindLeaderInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UnBindLeaderInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UnBindLeaderInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnBindLeaderInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindLeaderInstanceResponseBody) SetData(v bool) *UnBindLeaderInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) SetErrorCode(v string) *UnBindLeaderInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) SetErrorMessage(v string) *UnBindLeaderInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) SetHttpStatusCode(v string) *UnBindLeaderInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) SetRequestId(v string) *UnBindLeaderInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) SetSuccess(v bool) *UnBindLeaderInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UnBindLeaderInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
