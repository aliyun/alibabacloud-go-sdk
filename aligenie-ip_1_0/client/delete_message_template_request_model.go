// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *DeleteMessageTemplateRequest
	GetTemplateId() *int64
}

type DeleteMessageTemplateRequest struct {
	// example:
	//
	// 234
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMessageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteMessageTemplateRequest) SetTemplateId(v int64) *DeleteMessageTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteMessageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
