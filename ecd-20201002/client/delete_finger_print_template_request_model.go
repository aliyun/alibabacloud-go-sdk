// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFingerPrintTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteFingerPrintTemplateRequest
	GetClientId() *string
	SetClientToken(v string) *DeleteFingerPrintTemplateRequest
	GetClientToken() *string
	SetIndex(v int32) *DeleteFingerPrintTemplateRequest
	GetIndex() *int32
	SetLoginToken(v string) *DeleteFingerPrintTemplateRequest
	GetLoginToken() *string
	SetRegionId(v string) *DeleteFingerPrintTemplateRequest
	GetRegionId() *string
	SetSessionId(v string) *DeleteFingerPrintTemplateRequest
	GetSessionId() *string
}

type DeleteFingerPrintTemplateRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 59e86b39-ccac-4dfa-93d7-1f724052****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client token to ensure the idempotence of the request. You can use the client to generate the value, but you ensure sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 40401e62-5caf-4508-8de7-bf98af12****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The index.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1c0436c721786529914f16516396228454fa6284c9b80f9917f25ebbec2aa30c10343e3f6f9aff64500ce13808aef****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6df06330-3b75-4768-b334-41a73a64****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteFingerPrintTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteFingerPrintTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFingerPrintTemplateRequest) GetIndex() *int32 {
	return s.Index
}

func (s *DeleteFingerPrintTemplateRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DeleteFingerPrintTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFingerPrintTemplateRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteFingerPrintTemplateRequest) SetClientId(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetClientToken(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetIndex(v int32) *DeleteFingerPrintTemplateRequest {
	s.Index = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetLoginToken(v string) *DeleteFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetRegionId(v string) *DeleteFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetSessionId(v string) *DeleteFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) Validate() error {
	return dara.Validate(s)
}
