// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTemplateMCPRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopTemplateMCPRequest struct {
}

func (s StopTemplateMCPRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTemplateMCPRequest) GoString() string {
	return s.String()
}

func (s *StopTemplateMCPRequest) Validate() error {
	return dara.Validate(s)
}
