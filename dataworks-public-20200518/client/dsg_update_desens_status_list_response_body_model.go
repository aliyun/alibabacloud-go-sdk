// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUpdateDesensStatusListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgUpdateDesensStatusListResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgUpdateDesensStatusListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUpdateDesensStatusListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUpdateDesensStatusListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgUpdateDesensStatusListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUpdateDesensStatusListResponseBody
	GetSuccess() *bool
}

type DsgUpdateDesensStatusListResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 1010040007
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgUpdateDesensStatusListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUpdateDesensStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUpdateDesensStatusListResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgUpdateDesensStatusListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUpdateDesensStatusListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUpdateDesensStatusListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUpdateDesensStatusListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUpdateDesensStatusListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUpdateDesensStatusListResponseBody) SetData(v bool) *DsgUpdateDesensStatusListResponseBody {
	s.Data = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) SetErrorCode(v string) *DsgUpdateDesensStatusListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) SetErrorMessage(v string) *DsgUpdateDesensStatusListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) SetHttpStatusCode(v int32) *DsgUpdateDesensStatusListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) SetRequestId(v string) *DsgUpdateDesensStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) SetSuccess(v bool) *DsgUpdateDesensStatusListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponseBody) Validate() error {
	return dara.Validate(s)
}
