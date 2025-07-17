// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRetcodeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteRetcodeAppResponseBody
	GetCode() *int32
	SetData(v string) *DeleteRetcodeAppResponseBody
	GetData() *string
	SetMessage(v string) *DeleteRetcodeAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRetcodeAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRetcodeAppResponseBody
	GetSuccess() *bool
}

type DeleteRetcodeAppResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the Browser Monitoring task was deleted. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Additional information. The value description is as follows:
	//
	// - If the request is normal, return success.
	//
	// - If the request is abnormal, return specific abnormal information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01FF8DD9-A09C-47A1-895A-B6E321BE77B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful:
	//
	// - `true`: The operation was successful
	//
	// - `false`: The operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRetcodeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteRetcodeAppResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteRetcodeAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRetcodeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRetcodeAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRetcodeAppResponseBody) SetCode(v int32) *DeleteRetcodeAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetData(v string) *DeleteRetcodeAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetMessage(v string) *DeleteRetcodeAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetRequestId(v string) *DeleteRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetSuccess(v bool) *DeleteRetcodeAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) Validate() error {
	return dara.Validate(s)
}
