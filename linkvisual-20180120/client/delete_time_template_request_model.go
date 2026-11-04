// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTimeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteTimeTemplateRequest
	GetTemplateId() *string
}

type DeleteTimeTemplateRequest struct {
	// The ID of the time template. You can call the [QueryTimeTemplate](https://help.aliyun.com/document_detail/107890.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36f06322442e4e3f85a5ae8862c9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTimeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTimeTemplateRequest) SetTemplateId(v string) *DeleteTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteTimeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
