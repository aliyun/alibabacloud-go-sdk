// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoryDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetCategoryDetailResponseBody
	GetData() *string
	SetErrCode(v string) *GetCategoryDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetCategoryDetailResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetCategoryDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCategoryDetailResponseBody
	GetSuccess() *bool
}

type GetCategoryDetailResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {\\"taskId\\": 1699}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCategoryDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCategoryDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoryDetailResponseBody) GetData() *string {
	return s.Data
}

func (s *GetCategoryDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetCategoryDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetCategoryDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCategoryDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCategoryDetailResponseBody) SetData(v string) *GetCategoryDetailResponseBody {
	s.Data = &v
	return s
}

func (s *GetCategoryDetailResponseBody) SetErrCode(v string) *GetCategoryDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetCategoryDetailResponseBody) SetErrMessage(v string) *GetCategoryDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetCategoryDetailResponseBody) SetRequestId(v string) *GetCategoryDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoryDetailResponseBody) SetSuccess(v bool) *GetCategoryDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetCategoryDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
