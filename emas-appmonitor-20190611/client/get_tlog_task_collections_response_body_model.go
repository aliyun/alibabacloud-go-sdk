// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *GetTlogTaskCollectionsResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *GetTlogTaskCollectionsResponseBody
	GetMessage() *string
	SetModel(v interface{}) *GetTlogTaskCollectionsResponseBody
	GetModel() interface{}
	SetRequestId(v string) *GetTlogTaskCollectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTlogTaskCollectionsResponseBody
	GetSuccess() *bool
}

type GetTlogTaskCollectionsResponseBody struct {
	// example:
	//
	// 500
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// successful
	Message *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTlogTaskCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlogTaskCollectionsResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *GetTlogTaskCollectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTlogTaskCollectionsResponseBody) GetModel() interface{} {
	return s.Model
}

func (s *GetTlogTaskCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlogTaskCollectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTlogTaskCollectionsResponseBody) SetErrorCode(v int64) *GetTlogTaskCollectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTlogTaskCollectionsResponseBody) SetMessage(v string) *GetTlogTaskCollectionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTlogTaskCollectionsResponseBody) SetModel(v interface{}) *GetTlogTaskCollectionsResponseBody {
	s.Model = v
	return s
}

func (s *GetTlogTaskCollectionsResponseBody) SetRequestId(v string) *GetTlogTaskCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlogTaskCollectionsResponseBody) SetSuccess(v bool) *GetTlogTaskCollectionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTlogTaskCollectionsResponseBody) Validate() error {
	return dara.Validate(s)
}
