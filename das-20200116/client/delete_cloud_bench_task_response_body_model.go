// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudBenchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCloudBenchTaskResponseBody
	GetCode() *string
	SetData(v string) *DeleteCloudBenchTaskResponseBody
	GetData() *string
	SetMessage(v string) *DeleteCloudBenchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCloudBenchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteCloudBenchTaskResponseBody
	GetSuccess() *string
}

type DeleteCloudBenchTaskResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCloudBenchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudBenchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCloudBenchTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCloudBenchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCloudBenchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudBenchTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteCloudBenchTaskResponseBody) SetCode(v string) *DeleteCloudBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetData(v string) *DeleteCloudBenchTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetMessage(v string) *DeleteCloudBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetRequestId(v string) *DeleteCloudBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetSuccess(v string) *DeleteCloudBenchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
