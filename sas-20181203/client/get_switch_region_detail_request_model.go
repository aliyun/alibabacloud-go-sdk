// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwitchRegionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetSwitchRegionDetailRequest
	GetLang() *string
	SetType(v string) *GetSwitchRegionDetailRequest
	GetType() *string
}

type GetSwitchRegionDetailRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the migration operation. Valid values:
	//
	// 	- **sg_switch**: the migration of a server from a region in the Chinese mainland to the Singapore region.
	//
	// 	- **sls_meta_version_switch_stage_1**: the upgrade of log dictionaries.
	//
	// example:
	//
	// sg_switch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSwitchRegionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSwitchRegionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSwitchRegionDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSwitchRegionDetailRequest) GetType() *string {
	return s.Type
}

func (s *GetSwitchRegionDetailRequest) SetLang(v string) *GetSwitchRegionDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetSwitchRegionDetailRequest) SetType(v string) *GetSwitchRegionDetailRequest {
	s.Type = &v
	return s
}

func (s *GetSwitchRegionDetailRequest) Validate() error {
	return dara.Validate(s)
}
