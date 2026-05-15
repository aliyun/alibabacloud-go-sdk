// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartCallV2ResponseBody
	GetCode() *string
	SetMessage(v string) *StartCallV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartCallV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartCallV2ResponseBody
	GetSuccess() *bool
}

type StartCallV2ResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartCallV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartCallV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCallV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartCallV2ResponseBody) SetCode(v string) *StartCallV2ResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallV2ResponseBody) SetMessage(v string) *StartCallV2ResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallV2ResponseBody) SetRequestId(v string) *StartCallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallV2ResponseBody) SetSuccess(v bool) *StartCallV2ResponseBody {
	s.Success = &v
	return s
}

func (s *StartCallV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
