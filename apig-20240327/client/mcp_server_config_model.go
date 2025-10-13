// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMcpServerConfig interface {
	dara.Model
	String() string
	GoString() string
}

type McpServerConfig struct {
}

func (s McpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s McpServerConfig) GoString() string {
	return s.String()
}

func (s *McpServerConfig) Validate() error {
	return dara.Validate(s)
}
