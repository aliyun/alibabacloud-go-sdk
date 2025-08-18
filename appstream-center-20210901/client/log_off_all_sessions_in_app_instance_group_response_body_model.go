// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogOffAllSessionsInAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody
	GetCode() *string
	SetMessage(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody
	GetRequestId() *string
}

type LogOffAllSessionsInAppInstanceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetCode(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetMessage(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetRequestId(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
