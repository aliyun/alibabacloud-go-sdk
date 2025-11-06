// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppViewTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *ModifyAppViewTemplateResponseBody
	GetTemplateId() *string
}

type ModifyAppViewTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ac7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyAppViewTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppViewTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppViewTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyAppViewTemplateResponseBody) SetRequestId(v string) *ModifyAppViewTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppViewTemplateResponseBody) SetTemplateId(v string) *ModifyAppViewTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ModifyAppViewTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
