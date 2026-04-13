// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPublishedAgentRequest struct {
}

func (s GetPublishedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentRequest) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentRequest) Validate() error {
	return dara.Validate(s)
}
