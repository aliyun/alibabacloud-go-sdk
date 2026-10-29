// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsGetRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SdnsGetRuleCategoryRequest
	GetId() *string
}

type SdnsGetRuleCategoryRequest struct {
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SdnsGetRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryRequest) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryRequest) SetId(v string) *SdnsGetRuleCategoryRequest {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
