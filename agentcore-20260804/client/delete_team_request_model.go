// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTeamRequest
	GetClientToken() *string
}

type DeleteTeamRequest struct {
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTeamRequest) SetClientToken(v string) *DeleteTeamRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTeamRequest) Validate() error {
	return dara.Validate(s)
}
