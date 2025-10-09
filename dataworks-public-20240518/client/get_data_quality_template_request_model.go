// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDataQualityTemplateRequest
	GetId() *string
}

type GetDataQualityTemplateRequest struct {
	// The data quality rule template ID.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityTemplateRequest) GetId() *string {
	return s.Id
}

func (s *GetDataQualityTemplateRequest) SetId(v string) *GetDataQualityTemplateRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
