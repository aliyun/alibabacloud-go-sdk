// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpProductToolsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListMcpProductToolsRequest struct {
}

func (s ListMcpProductToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpProductToolsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpProductToolsRequest) Validate() error {
	return dara.Validate(s)
}
