// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPortalId(v string) *ListPortalMenusRequest
	GetDataPortalId() *string
	SetUserId(v string) *ListPortalMenusRequest
	GetUserId() *string
}

type ListPortalMenusRequest struct {
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The user ID in the Quick BI. When passed in, the list displays only the menus that the user has permissions on.
	//
	// example:
	//
	// 1234567***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPortalMenusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenusRequest) GoString() string {
	return s.String()
}

func (s *ListPortalMenusRequest) GetDataPortalId() *string {
	return s.DataPortalId
}

func (s *ListPortalMenusRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListPortalMenusRequest) SetDataPortalId(v string) *ListPortalMenusRequest {
	s.DataPortalId = &v
	return s
}

func (s *ListPortalMenusRequest) SetUserId(v string) *ListPortalMenusRequest {
	s.UserId = &v
	return s
}

func (s *ListPortalMenusRequest) Validate() error {
	return dara.Validate(s)
}
