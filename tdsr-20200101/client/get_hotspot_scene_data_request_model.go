// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotSceneDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetHotspotSceneDataRequest
	GetDomain() *string
	SetEnabled(v bool) *GetHotspotSceneDataRequest
	GetEnabled() *bool
	SetPreviewToken(v string) *GetHotspotSceneDataRequest
	GetPreviewToken() *string
	SetType(v int64) *GetHotspotSceneDataRequest
	GetType() *int64
}

type GetHotspotSceneDataRequest struct {
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
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

func (s GetHotspotSceneDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotSceneDataRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetHotspotSceneDataRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetHotspotSceneDataRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetHotspotSceneDataRequest) GetType() *int64 {
	return s.Type
}

func (s *GetHotspotSceneDataRequest) SetDomain(v string) *GetHotspotSceneDataRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetEnabled(v bool) *GetHotspotSceneDataRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetPreviewToken(v string) *GetHotspotSceneDataRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetType(v int64) *GetHotspotSceneDataRequest {
	s.Type = &v
	return s
}

func (s *GetHotspotSceneDataRequest) Validate() error {
	return dara.Validate(s)
}
