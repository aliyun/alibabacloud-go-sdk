// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetScenePreviewInfoRequest
	GetDomain() *string
	SetEnabled(v bool) *GetScenePreviewInfoRequest
	GetEnabled() *bool
	SetModelToken(v string) *GetScenePreviewInfoRequest
	GetModelToken() *string
}

type GetScenePreviewInfoRequest struct {
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
	// A.e.RQJRPYGIJJQP****
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
}

func (s GetScenePreviewInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetScenePreviewInfoRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetScenePreviewInfoRequest) GetModelToken() *string {
	return s.ModelToken
}

func (s *GetScenePreviewInfoRequest) SetDomain(v string) *GetScenePreviewInfoRequest {
	s.Domain = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetEnabled(v bool) *GetScenePreviewInfoRequest {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetModelToken(v string) *GetScenePreviewInfoRequest {
	s.ModelToken = &v
	return s
}

func (s *GetScenePreviewInfoRequest) Validate() error {
	return dara.Validate(s)
}
