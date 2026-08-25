// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChallengeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetChallengeRequest struct {
}

func (s GetChallengeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChallengeRequest) GoString() string {
	return s.String()
}

func (s *GetChallengeRequest) Validate() error {
	return dara.Validate(s)
}
