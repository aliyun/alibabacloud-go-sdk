// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetHotspotConfigRequest
	GetDomain() *string
	SetEnabled(v bool) *GetHotspotConfigRequest
	GetEnabled() *bool
	SetPreviewToken(v string) *GetHotspotConfigRequest
	GetPreviewToken() *string
	SetType(v int64) *GetHotspotConfigRequest
	GetType() *int64
}

type GetHotspotConfigRequest struct {
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true/false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 344794c32937474a9c59eb130936****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetHotspotConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetHotspotConfigRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetHotspotConfigRequest) GetType() *int64 {
	return s.Type
}

func (s *GetHotspotConfigRequest) SetDomain(v string) *GetHotspotConfigRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotConfigRequest) SetEnabled(v bool) *GetHotspotConfigRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotConfigRequest) SetPreviewToken(v string) *GetHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotConfigRequest) SetType(v int64) *GetHotspotConfigRequest {
	s.Type = &v
	return s
}

func (s *GetHotspotConfigRequest) Validate() error {
	return dara.Validate(s)
}
