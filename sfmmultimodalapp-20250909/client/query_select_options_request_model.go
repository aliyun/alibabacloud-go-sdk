// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySelectOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *QuerySelectOptionsRequest
	GetWorkspaceId() *string
}

type QuerySelectOptionsRequest struct {
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QuerySelectOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySelectOptionsRequest) GoString() string {
	return s.String()
}

func (s *QuerySelectOptionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QuerySelectOptionsRequest) SetWorkspaceId(v string) *QuerySelectOptionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QuerySelectOptionsRequest) Validate() error {
	return dara.Validate(s)
}
