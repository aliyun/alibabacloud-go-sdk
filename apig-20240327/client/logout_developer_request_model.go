// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogoutDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
}

type LogoutDeveloperRequest struct {
}

func (s LogoutDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s LogoutDeveloperRequest) GoString() string {
	return s.String()
}

func (s *LogoutDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
