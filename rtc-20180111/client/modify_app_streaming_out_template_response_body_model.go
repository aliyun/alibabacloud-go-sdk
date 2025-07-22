// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppStreamingOutTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppStreamingOutTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *ModifyAppStreamingOutTemplateResponseBody
	GetTemplateId() *string
}

type ModifyAppStreamingOutTemplateResponseBody struct {
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

func (s ModifyAppStreamingOutTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppStreamingOutTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyAppStreamingOutTemplateResponseBody) SetRequestId(v string) *ModifyAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponseBody) SetTemplateId(v string) *ModifyAppStreamingOutTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
