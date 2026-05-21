// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindLeaderInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *BindLeaderInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *BindLeaderInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BindLeaderInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *BindLeaderInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *BindLeaderInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindLeaderInstanceResponseBody
	GetSuccess() *bool
}

type BindLeaderInstanceResponseBody struct {
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
	// 6FF2EDFD-3154-5C3E-8CFA-8F7E366BCF9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindLeaderInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindLeaderInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BindLeaderInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindLeaderInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BindLeaderInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BindLeaderInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *BindLeaderInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindLeaderInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindLeaderInstanceResponseBody) SetData(v bool) *BindLeaderInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) SetErrorCode(v string) *BindLeaderInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) SetErrorMessage(v string) *BindLeaderInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) SetHttpStatusCode(v string) *BindLeaderInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) SetRequestId(v string) *BindLeaderInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) SetSuccess(v bool) *BindLeaderInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *BindLeaderInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
