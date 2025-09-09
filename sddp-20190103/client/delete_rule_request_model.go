// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DeleteRuleRequest
	GetFeatureType() *int32
	SetId(v int64) *DeleteRuleRequest
	GetId() *int64
	SetLang(v string) *DeleteRuleRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteRuleRequest
	GetSourceIp() *string
}

type DeleteRuleRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 122300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values: **zh*	- and **en**. The value zh indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DeleteRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteRuleRequest) SetFeatureType(v int32) *DeleteRuleRequest {
	s.FeatureType = &v
	return s
}

func (s *DeleteRuleRequest) SetId(v int64) *DeleteRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetSourceIp(v string) *DeleteRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
