// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoForkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ReposForkCreateCmd) *CreateRepoForkRequest
	GetBody() *ReposForkCreateCmd
}

type CreateRepoForkRequest struct {
	Body *ReposForkCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoForkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoForkRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoForkRequest) GetBody() *ReposForkCreateCmd {
	return s.Body
}

func (s *CreateRepoForkRequest) SetBody(v *ReposForkCreateCmd) *CreateRepoForkRequest {
	s.Body = v
	return s
}

func (s *CreateRepoForkRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
