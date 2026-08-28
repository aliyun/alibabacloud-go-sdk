// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListPluginRepositoriesRequest struct {
}

func (s ListPluginRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListPluginRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
