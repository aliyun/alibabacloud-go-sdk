// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCursorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CursorResponseBody
	GetCode() *int32
	SetCursor(v string) *CursorResponseBody
	GetCursor() *string
	SetMessage(v string) *CursorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CursorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CursorResponseBody
	GetSuccess() *bool
}

type CursorResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Cursor is used as an input parameter for data export in the [BatchExport](https://help.aliyun.com/document_detail/2329847.html) operation.
	//
	// example:
	//
	// v2.5eyJiYXRjaGVzIjoxLCJidWNrZXRCeXRlcyI6IndBPT0iLCJidWNrZXRzIjo0LCJjdXJyZW50QnVja2V0IjotNjQsImN1cnJlbnRUYXJnZXRUaW1lU2xvdCI6MTY3ODc1MjAwMCwiZW5kVGltZSI6MTY3ODc4MjQxOTAwMCwiZXhwb3J0RW5kVGltZSI6MTY3ODc4MjQxOTAwMCwiZXhwb3J0U3RhcnRUaW1lIjoxNjc4NzgxODE5MDAwLCJleHByZXNzUmFuZ2UiOmZhbHNlLCJoYXNOZXh0Ijp0cnVlLCJpbmRleCI6MCwibGF0ZXN0TG9nVGltZSI6MCwibWF0Y2hlcnMiOnsiY2hhaW4iOlt7ImxhYmVsIjoidXNlcklkIiwib3BlcmF0b3IiOiJFUVVBTFMiLCJ2YWx1ZSI6IjEyNzA2NzY2Nzk1NDY3MDQifV0sImxvY2tlZCI6dHJ1ZX0sIm1ldHJpYyI6IlNwbGl0cndQcm94eU1heFJlc3BvbnNlU2l6ZSIsIm1ldHJpY1R5cGUiOiJNRVRSSUMiLCJuYW1lc3BhY2UiOiJhY3Nfa3ZzdG9yZV9leCIsIm5leHRQa0FkYXB0ZXIiOnsiZGltIjoiVjowXG5EOmB1c2VySWRgPTEyNzA2NzY2Nzk1NDY3MDQsYGluc3RhbmNlSWRgPXItcmo5ZjlzMTlsc3V1MXd1bnVyLGBub2RlSWRgPXItcmo5ZjlzMTlsc3V1MXd1bnVyLXByb3h5LTIsXG4iLCJtZXRhIjoiXG5WOjBcbk06YWNzX2t2c3RvcmVfZXgvU3BsaXRyd1Byb3h5TWF4UmVzcG9uc2VTaXplXG5XOjYwXG5COjRcbkk6LTFcblQ6MFxuQzpgQXZlcmFnZWAsYE1heGltdW1gLGBfX2NvdW50X19gLGBfX3RzX19gXG4iLCJyZCI6InN1YkFMU0RwWXY2K0t6aENQQUFBWkErNUFKMEpjbGErRGd2V0hFNWxDSHMvbGtqR2FXYTFJTkVwdFE9PSIsInRhZyI6IiJ9LCJvZmZzZXQiOjAsIm9mZnNldERpZ2l0Ijo0NTU0NTczNDQyMTc4NDIxMjIsInN0YXJ0VGltZSI6MTY3ODc4MTgxOTAwMCwic3RlcCI6LTEsInRhZ01hdGNoZXJzIjp7ImNoYWluIjpbXSwibG9ja2VkIjp0cnVlfSwidGFyZ2V0VGltZVNsb3RzIjpbMTY3ODY2NTYwMCwxNjc4NzUyMDAwXSwidXVpZCI6ImQwMmFhZmY1LWU3ZGQtNDUyYy0***********
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 915C2D7A-E6A4-17C3-8E13-4DBDD61E7919
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CursorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CursorResponseBody) GoString() string {
	return s.String()
}

func (s *CursorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CursorResponseBody) GetCursor() *string {
	return s.Cursor
}

func (s *CursorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CursorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CursorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CursorResponseBody) SetCode(v int32) *CursorResponseBody {
	s.Code = &v
	return s
}

func (s *CursorResponseBody) SetCursor(v string) *CursorResponseBody {
	s.Cursor = &v
	return s
}

func (s *CursorResponseBody) SetMessage(v string) *CursorResponseBody {
	s.Message = &v
	return s
}

func (s *CursorResponseBody) SetRequestId(v string) *CursorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CursorResponseBody) SetSuccess(v bool) *CursorResponseBody {
	s.Success = &v
	return s
}

func (s *CursorResponseBody) Validate() error {
	return dara.Validate(s)
}
