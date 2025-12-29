// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryDeviceStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryDeviceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDeviceStatusResponseBody
	GetRequestId() *string
	SetResult(v []map[string]*string) *QueryDeviceStatusResponseBody
	GetResult() []map[string]*string
}

type QueryDeviceStatusResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsgrefds
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryDeviceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDeviceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceStatusResponseBody) GetResult() []map[string]*string {
	return s.Result
}

func (s *QueryDeviceStatusResponseBody) SetCode(v int32) *QueryDeviceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetMessage(v string) *QueryDeviceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetRequestId(v string) *QueryDeviceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceStatusResponseBody) SetResult(v []map[string]*string) *QueryDeviceStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryDeviceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
