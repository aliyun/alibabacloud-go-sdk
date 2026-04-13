// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CommitServiceRequest struct {
}

func (s CommitServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitServiceRequest) GoString() string {
	return s.String()
}

func (s *CommitServiceRequest) Validate() error {
	return dara.Validate(s)
}
