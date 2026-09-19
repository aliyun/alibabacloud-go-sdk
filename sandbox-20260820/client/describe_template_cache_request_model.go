// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *DescribeTemplateCacheRequest
	GetTeamID() *string
}

type DescribeTemplateCacheRequest struct {
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s DescribeTemplateCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCacheRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *DescribeTemplateCacheRequest) SetTeamID(v string) *DescribeTemplateCacheRequest {
	s.TeamID = &v
	return s
}

func (s *DescribeTemplateCacheRequest) Validate() error {
	return dara.Validate(s)
}
