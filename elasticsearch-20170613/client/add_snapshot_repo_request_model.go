// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnapshotRepoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *AddSnapshotRepoRequest
	GetBody() *string
}

type AddSnapshotRepoRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSnapshotRepoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSnapshotRepoRequest) GoString() string {
	return s.String()
}

func (s *AddSnapshotRepoRequest) GetBody() *string {
	return s.Body
}

func (s *AddSnapshotRepoRequest) SetBody(v string) *AddSnapshotRepoRequest {
	s.Body = &v
	return s
}

func (s *AddSnapshotRepoRequest) Validate() error {
	return dara.Validate(s)
}
