// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteConfigTemplateRequest
	GetId() *int64
}

type DeleteConfigTemplateRequest struct {
	// The ID of the configuration template.
	//
	// example:
	//
	// 555
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteConfigTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteConfigTemplateRequest) SetId(v int64) *DeleteConfigTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteConfigTemplateRequest) Validate() error {
	return dara.Validate(s)
}
