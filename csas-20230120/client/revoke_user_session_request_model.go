// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIds(v string) *RevokeUserSessionRequest
	GetExternalIds() *string
	SetIdpId(v string) *RevokeUserSessionRequest
	GetIdpId() *string
}

type RevokeUserSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	ExternalIds *string `json:"ExternalIds,omitempty" xml:"ExternalIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// idp-cfg9vcrqylo39c39uxnw
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s RevokeUserSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserSessionRequest) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionRequest) GetExternalIds() *string {
	return s.ExternalIds
}

func (s *RevokeUserSessionRequest) GetIdpId() *string {
	return s.IdpId
}

func (s *RevokeUserSessionRequest) SetExternalIds(v string) *RevokeUserSessionRequest {
	s.ExternalIds = &v
	return s
}

func (s *RevokeUserSessionRequest) SetIdpId(v string) *RevokeUserSessionRequest {
	s.IdpId = &v
	return s
}

func (s *RevokeUserSessionRequest) Validate() error {
	return dara.Validate(s)
}
