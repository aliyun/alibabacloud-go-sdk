// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdvancedMonitorStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetAdvancedMonitorStateResponseBody
	GetCode() *string
	SetMessage(v string) *SetAdvancedMonitorStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetAdvancedMonitorStateResponseBody
	GetRequestId() *string
}

type SetAdvancedMonitorStateResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAdvancedMonitorStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAdvancedMonitorStateResponseBody) GoString() string {
	return s.String()
}

func (s *SetAdvancedMonitorStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetAdvancedMonitorStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetAdvancedMonitorStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAdvancedMonitorStateResponseBody) SetCode(v string) *SetAdvancedMonitorStateResponseBody {
	s.Code = &v
	return s
}

func (s *SetAdvancedMonitorStateResponseBody) SetMessage(v string) *SetAdvancedMonitorStateResponseBody {
	s.Message = &v
	return s
}

func (s *SetAdvancedMonitorStateResponseBody) SetRequestId(v string) *SetAdvancedMonitorStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAdvancedMonitorStateResponseBody) Validate() error {
	return dara.Validate(s)
}
