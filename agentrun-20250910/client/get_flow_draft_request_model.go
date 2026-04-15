// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowDraftRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFlowDraftRequest struct {
}

func (s GetFlowDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowDraftRequest) GoString() string {
	return s.String()
}

func (s *GetFlowDraftRequest) Validate() error {
	return dara.Validate(s)
}
