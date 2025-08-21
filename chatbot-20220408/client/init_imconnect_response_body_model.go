// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitIMConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitIMConnectResponseBody
	GetCode() *string
	SetData(v string) *InitIMConnectResponseBody
	GetData() *string
	SetMessage(v string) *InitIMConnectResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitIMConnectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitIMConnectResponseBody
	GetSuccess() *bool
}

type InitIMConnectResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {     "imDomain": "im.alimebot.com",     "appKey": "WDg2VfNv",     "token": "QUM4SndaY3VPMjhkQldDZUNOR0ZaTmZ5R3NBY0FKWHJ4OGc4dERZbEJzcjNIKzFiS1RyTjhXRUpBYmVpQlpsakprNDRFVkdxcy9HWVk2RXZvalU3bHhxRkJlc1NBUXZwdHFKOTE2UTNwamQ4b1U4N3dEbmhyRjc4R2hOQStvMnMrYkV2dlVpSHNvWC96SEVNZWRqMjBuMXdjNklpamJzaDNWYllnUldDZGhJPQ=="   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E6988CE6-41CF-1103-9BEC-2B20D26C0B52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitIMConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitIMConnectResponseBody) GoString() string {
	return s.String()
}

func (s *InitIMConnectResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitIMConnectResponseBody) GetData() *string {
	return s.Data
}

func (s *InitIMConnectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitIMConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitIMConnectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitIMConnectResponseBody) SetCode(v string) *InitIMConnectResponseBody {
	s.Code = &v
	return s
}

func (s *InitIMConnectResponseBody) SetData(v string) *InitIMConnectResponseBody {
	s.Data = &v
	return s
}

func (s *InitIMConnectResponseBody) SetMessage(v string) *InitIMConnectResponseBody {
	s.Message = &v
	return s
}

func (s *InitIMConnectResponseBody) SetRequestId(v string) *InitIMConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitIMConnectResponseBody) SetSuccess(v bool) *InitIMConnectResponseBody {
	s.Success = &v
	return s
}

func (s *InitIMConnectResponseBody) Validate() error {
	return dara.Validate(s)
}
