// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAppResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartAppResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *RestartAppResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartAppResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *RestartAppResourcesResponseBody
	GetSuccess() *string
}

type RestartAppResourcesResponseBody struct {
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

func (s RestartAppResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartAppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartAppResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartAppResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartAppResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartAppResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *RestartAppResourcesResponseBody) SetCode(v string) *RestartAppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *RestartAppResourcesResponseBody) SetMessage(v string) *RestartAppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *RestartAppResourcesResponseBody) SetRequestId(v string) *RestartAppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartAppResourcesResponseBody) SetSuccess(v string) *RestartAppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartAppResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
