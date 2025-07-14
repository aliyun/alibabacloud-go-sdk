// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenuAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPortalId(v string) *ListPortalMenuAuthorizationRequest
	GetDataPortalId() *string
}

type ListPortalMenuAuthorizationRequest struct {
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
}

func (s ListPortalMenuAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenuAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationRequest) GetDataPortalId() *string {
	return s.DataPortalId
}

func (s *ListPortalMenuAuthorizationRequest) SetDataPortalId(v string) *ListPortalMenuAuthorizationRequest {
	s.DataPortalId = &v
	return s
}

func (s *ListPortalMenuAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
