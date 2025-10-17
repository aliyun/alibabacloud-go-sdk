// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTemplateRequest
	GetRegionId() *string
}

type ListTemplateRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTemplateRequest) SetRegionId(v string) *ListTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateRequest) Validate() error {
	return dara.Validate(s)
}
