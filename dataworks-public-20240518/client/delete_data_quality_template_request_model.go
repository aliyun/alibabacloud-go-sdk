// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDataQualityTemplateRequest
	GetId() *string
}

type DeleteDataQualityTemplateRequest struct {
	// example:
	//
	// USER_DEFINED:2001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityTemplateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDataQualityTemplateRequest) SetId(v string) *DeleteDataQualityTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataQualityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
