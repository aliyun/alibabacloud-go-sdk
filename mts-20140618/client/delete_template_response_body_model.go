// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DeleteTemplateResponseBody
	GetTemplateId() *string
}

type DeleteTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2247541A-9F27-47EE-B6EC-484B5475****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the custom transcoding template that is deleted.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetTemplateId(v string) *DeleteTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
