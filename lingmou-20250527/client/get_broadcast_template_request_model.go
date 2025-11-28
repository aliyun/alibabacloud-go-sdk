// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBroadcastTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetBroadcastTemplateRequest
	GetTemplateId() *string
}

type GetBroadcastTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BS1b2WNnRMu4ouRzT4clY9Jhg
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s GetBroadcastTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBroadcastTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetBroadcastTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetBroadcastTemplateRequest) SetTemplateId(v string) *GetBroadcastTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetBroadcastTemplateRequest) Validate() error {
	return dara.Validate(s)
}
