// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppStreamingOutTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppStreamingOutTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateAppStreamingOutTemplateResponseBody
	GetTemplateId() *string
}

type CreateAppStreamingOutTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// bc5v****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateAppStreamingOutTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppStreamingOutTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateAppStreamingOutTemplateResponseBody) SetRequestId(v string) *CreateAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateResponseBody) SetTemplateId(v string) *CreateAppStreamingOutTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
