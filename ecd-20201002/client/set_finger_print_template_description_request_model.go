// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *SetFingerPrintTemplateDescriptionRequest
	GetClientId() *string
	SetClientToken(v string) *SetFingerPrintTemplateDescriptionRequest
	GetClientToken() *string
	SetDescription(v string) *SetFingerPrintTemplateDescriptionRequest
	GetDescription() *string
	SetIndex(v int32) *SetFingerPrintTemplateDescriptionRequest
	GetIndex() *int32
	SetLoginToken(v string) *SetFingerPrintTemplateDescriptionRequest
	GetLoginToken() *string
	SetRegionId(v string) *SetFingerPrintTemplateDescriptionRequest
	GetRegionId() *string
	SetSessionId(v string) *SetFingerPrintTemplateDescriptionRequest
	GetSessionId() *string
}

type SetFingerPrintTemplateDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0764064c-1609-4d3c-8cb7-ab8d3feg****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 40401e62-5caf-4508-8de7-bf98af12****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Finger 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v14e5a2404c495249f7541646535779667ea0b5d87754b5d2d2a3099bda774f3832e24756ef3e66eb574b1f3e99078****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d28520d4-da0b-4a97-981d-683db865****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetClientId() *string {
	return s.ClientId
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetIndex() *int32 {
	return s.Index
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetFingerPrintTemplateDescriptionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetDescription(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetIndex(v int32) *SetFingerPrintTemplateDescriptionRequest {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetLoginToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetRegionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetSessionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.SessionId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
