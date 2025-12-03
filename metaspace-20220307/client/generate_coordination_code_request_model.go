// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCoordinationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginRegionId(v string) *GenerateCoordinationCodeRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *GenerateCoordinationCodeRequest
	GetLoginToken() *string
	SetResourceId(v string) *GenerateCoordinationCodeRequest
	GetResourceId() *string
	SetResourceName(v string) *GenerateCoordinationCodeRequest
	GetResourceName() *string
	SetResourceRegionId(v string) *GenerateCoordinationCodeRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *GenerateCoordinationCodeRequest
	GetResourceType() *string
	SetSessionId(v string) *GenerateCoordinationCodeRequest
	GetSessionId() *string
}

type GenerateCoordinationCodeRequest struct {
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// ecd-68a7ddrt0******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// demo-desktop
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GenerateCoordinationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCoordinationCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *GenerateCoordinationCodeRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GenerateCoordinationCodeRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GenerateCoordinationCodeRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *GenerateCoordinationCodeRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *GenerateCoordinationCodeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GenerateCoordinationCodeRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateCoordinationCodeRequest) SetLoginRegionId(v string) *GenerateCoordinationCodeRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetLoginToken(v string) *GenerateCoordinationCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceId(v string) *GenerateCoordinationCodeRequest {
	s.ResourceId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceName(v string) *GenerateCoordinationCodeRequest {
	s.ResourceName = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceRegionId(v string) *GenerateCoordinationCodeRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetResourceType(v string) *GenerateCoordinationCodeRequest {
	s.ResourceType = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetSessionId(v string) *GenerateCoordinationCodeRequest {
	s.SessionId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) Validate() error {
	return dara.Validate(s)
}
