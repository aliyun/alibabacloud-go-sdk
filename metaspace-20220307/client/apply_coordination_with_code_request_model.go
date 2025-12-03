// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationWithCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinationCode(v string) *ApplyCoordinationWithCodeRequest
	GetCoordinationCode() *string
	SetLoginRegionId(v string) *ApplyCoordinationWithCodeRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ApplyCoordinationWithCodeRequest
	GetLoginToken() *string
	SetSessionId(v string) *ApplyCoordinationWithCodeRequest
	GetSessionId() *string
	SetUuid(v string) *ApplyCoordinationWithCodeRequest
	GetUuid() *string
}

type ApplyCoordinationWithCodeRequest struct {
	// example:
	//
	// PA3MU***
	CoordinationCode *string `json:"CoordinationCode,omitempty" xml:"CoordinationCode,omitempty"`
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
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// A8B35215993FBF283F28D61******
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyCoordinationWithCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationWithCodeRequest) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeRequest) GetCoordinationCode() *string {
	return s.CoordinationCode
}

func (s *ApplyCoordinationWithCodeRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ApplyCoordinationWithCodeRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ApplyCoordinationWithCodeRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ApplyCoordinationWithCodeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ApplyCoordinationWithCodeRequest) SetCoordinationCode(v string) *ApplyCoordinationWithCodeRequest {
	s.CoordinationCode = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetLoginRegionId(v string) *ApplyCoordinationWithCodeRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetLoginToken(v string) *ApplyCoordinationWithCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetSessionId(v string) *ApplyCoordinationWithCodeRequest {
	s.SessionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) SetUuid(v string) *ApplyCoordinationWithCodeRequest {
	s.Uuid = &v
	return s
}

func (s *ApplyCoordinationWithCodeRequest) Validate() error {
	return dara.Validate(s)
}
