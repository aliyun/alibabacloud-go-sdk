// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTemplateResponseBody
	GetCode() *string
	SetMsg(v string) *DeleteTemplateResponseBody
	GetMsg() *string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTemplateResponseBody
	GetSuccess() *bool
}

type DeleteTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTemplateResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTemplateResponseBody) SetCode(v string) *DeleteTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetMsg(v string) *DeleteTemplateResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetSuccess(v bool) *DeleteTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
