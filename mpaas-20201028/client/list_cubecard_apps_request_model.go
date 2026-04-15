// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCubecardAppsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListCubecardAppsRequest struct {
}

func (s ListCubecardAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsRequest) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsRequest) Validate() error {
	return dara.Validate(s)
}
