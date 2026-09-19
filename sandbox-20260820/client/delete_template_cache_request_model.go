// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *DeleteTemplateCacheRequest
	GetTeamID() *string
}

type DeleteTemplateCacheRequest struct {
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s DeleteTemplateCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateCacheRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *DeleteTemplateCacheRequest) SetTeamID(v string) *DeleteTemplateCacheRequest {
	s.TeamID = &v
	return s
}

func (s *DeleteTemplateCacheRequest) Validate() error {
	return dara.Validate(s)
}
