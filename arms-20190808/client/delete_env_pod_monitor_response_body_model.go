// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvPodMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEnvPodMonitorResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEnvPodMonitorResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEnvPodMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvPodMonitorResponseBody
	GetRequestId() *string
}

type DeleteEnvPodMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 626037F5-FDEB-45B0-804C-B3C92797****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnvPodMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvPodMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvPodMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEnvPodMonitorResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEnvPodMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvPodMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvPodMonitorResponseBody) SetCode(v int32) *DeleteEnvPodMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvPodMonitorResponseBody) SetData(v string) *DeleteEnvPodMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnvPodMonitorResponseBody) SetMessage(v string) *DeleteEnvPodMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvPodMonitorResponseBody) SetRequestId(v string) *DeleteEnvPodMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvPodMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
