// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetAppResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *ResetAppResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetAppResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ResetAppResourcesResponseBody
	GetSuccess() *string
}

type ResetAppResourcesResponseBody struct {
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

func (s ResetAppResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetAppResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetAppResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAppResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ResetAppResourcesResponseBody) SetCode(v string) *ResetAppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ResetAppResourcesResponseBody) SetMessage(v string) *ResetAppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAppResourcesResponseBody) SetRequestId(v string) *ResetAppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAppResourcesResponseBody) SetSuccess(v string) *ResetAppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ResetAppResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
