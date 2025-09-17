// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAppResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAppResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *StartAppResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAppResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartAppResourcesResponseBody
	GetSuccess() *string
}

type StartAppResourcesResponseBody struct {
	// example:
	//
	// InvalidAppInstanceGroup.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The app instance group is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartAppResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *StartAppResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAppResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAppResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAppResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartAppResourcesResponseBody) SetCode(v string) *StartAppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *StartAppResourcesResponseBody) SetMessage(v string) *StartAppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *StartAppResourcesResponseBody) SetRequestId(v string) *StartAppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAppResourcesResponseBody) SetSuccess(v string) *StartAppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *StartAppResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
