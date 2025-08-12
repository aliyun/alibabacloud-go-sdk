// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitoringConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PutMonitoringConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *PutMonitoringConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutMonitoringConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutMonitoringConfigResponseBody
	GetSuccess() *bool
}

type PutMonitoringConfigResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Specified parameter EnableInstallAgentNewECS is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 109C8095-6FAD-4DBB-B013-6ED18CE4C0B1
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

func (s PutMonitoringConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutMonitoringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PutMonitoringConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutMonitoringConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutMonitoringConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutMonitoringConfigResponseBody) SetCode(v int32) *PutMonitoringConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetMessage(v string) *PutMonitoringConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetRequestId(v string) *PutMonitoringConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetSuccess(v bool) *PutMonitoringConfigResponseBody {
	s.Success = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
