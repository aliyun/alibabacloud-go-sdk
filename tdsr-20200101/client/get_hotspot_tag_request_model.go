// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetHotspotTagRequest
	GetDomain() *string
	SetEnabled(v bool) *GetHotspotTagRequest
	GetEnabled() *bool
	SetPreviewToken(v string) *GetHotspotTagRequest
	GetPreviewToken() *string
	SetSubSceneUuid(v string) *GetHotspotTagRequest
	GetSubSceneUuid() *string
	SetType(v string) *GetHotspotTagRequest
	GetType() *string
}

type GetHotspotTagRequest struct {
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true/false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb130936****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// example:
	//
	// 123456
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotTagRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetHotspotTagRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetHotspotTagRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetHotspotTagRequest) GetSubSceneUuid() *string {
	return s.SubSceneUuid
}

func (s *GetHotspotTagRequest) GetType() *string {
	return s.Type
}

func (s *GetHotspotTagRequest) SetDomain(v string) *GetHotspotTagRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotTagRequest) SetEnabled(v bool) *GetHotspotTagRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotTagRequest) SetPreviewToken(v string) *GetHotspotTagRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotTagRequest) SetSubSceneUuid(v string) *GetHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *GetHotspotTagRequest) SetType(v string) *GetHotspotTagRequest {
	s.Type = &v
	return s
}

func (s *GetHotspotTagRequest) Validate() error {
	return dara.Validate(s)
}
