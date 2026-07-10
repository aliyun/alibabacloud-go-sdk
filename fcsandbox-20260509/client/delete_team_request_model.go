// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteTeamRequest struct {
}

func (s DeleteTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeamRequest) Validate() error {
	return dara.Validate(s)
}
