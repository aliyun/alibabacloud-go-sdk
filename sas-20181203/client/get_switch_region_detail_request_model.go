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
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The service switchover type. Valid values:
	//
	// - **sg_switch**: The server connection is migrated from China to Singapore.
	//
	// - **sls_meta_version_switch_stage_1**: The log analysis delivery field upgrade switchover.
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
