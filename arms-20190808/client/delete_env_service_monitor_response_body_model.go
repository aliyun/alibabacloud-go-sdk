// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvServiceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEnvServiceMonitorResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEnvServiceMonitorResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEnvServiceMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvServiceMonitorResponseBody
	GetRequestId() *string
}

type DeleteEnvServiceMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnvServiceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvServiceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvServiceMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEnvServiceMonitorResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEnvServiceMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvServiceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvServiceMonitorResponseBody) SetCode(v int32) *DeleteEnvServiceMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvServiceMonitorResponseBody) SetData(v string) *DeleteEnvServiceMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnvServiceMonitorResponseBody) SetMessage(v string) *DeleteEnvServiceMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvServiceMonitorResponseBody) SetRequestId(v string) *DeleteEnvServiceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvServiceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
