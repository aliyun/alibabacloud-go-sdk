// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DeleteTemplatesRequest
	GetBizType() *string
	SetTemplateIds(v []*string) *DeleteTemplatesRequest
	GetTemplateIds() []*string
}

type DeleteTemplatesRequest struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The IDs of the templates that you want to delete.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s DeleteTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesRequest) GetBizType() *string {
	return s.BizType
}

func (s *DeleteTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DeleteTemplatesRequest) SetBizType(v string) *DeleteTemplatesRequest {
	s.BizType = &v
	return s
}

func (s *DeleteTemplatesRequest) SetTemplateIds(v []*string) *DeleteTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *DeleteTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
