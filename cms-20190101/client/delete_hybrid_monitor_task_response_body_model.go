// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHybridMonitorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHybridMonitorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHybridMonitorTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteHybridMonitorTaskResponseBody
	GetSuccess() *string
}

type DeleteHybridMonitorTaskResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D6B8103F-7BAD-5FEB-945F-71D991AB8834
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHybridMonitorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHybridMonitorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHybridMonitorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridMonitorTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteHybridMonitorTaskResponseBody) SetCode(v string) *DeleteHybridMonitorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHybridMonitorTaskResponseBody) SetMessage(v string) *DeleteHybridMonitorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHybridMonitorTaskResponseBody) SetRequestId(v string) *DeleteHybridMonitorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridMonitorTaskResponseBody) SetSuccess(v string) *DeleteHybridMonitorTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHybridMonitorTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
