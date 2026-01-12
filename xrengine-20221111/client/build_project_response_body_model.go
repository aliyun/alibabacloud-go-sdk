// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuildProjectResponseBody
	GetCode() *string
	SetMessage(v string) *BuildProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *BuildProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuildProjectResponseBody
	GetSuccess() *bool
}

type BuildProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuildProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildProjectResponseBody) GoString() string {
	return s.String()
}

func (s *BuildProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuildProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BuildProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildProjectResponseBody) SetCode(v string) *BuildProjectResponseBody {
	s.Code = &v
	return s
}

func (s *BuildProjectResponseBody) SetMessage(v string) *BuildProjectResponseBody {
	s.Message = &v
	return s
}

func (s *BuildProjectResponseBody) SetRequestId(v string) *BuildProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildProjectResponseBody) SetSuccess(v bool) *BuildProjectResponseBody {
	s.Success = &v
	return s
}

func (s *BuildProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
