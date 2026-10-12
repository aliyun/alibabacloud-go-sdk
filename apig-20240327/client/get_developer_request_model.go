// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDeveloperRequest struct {
}

func (s GetDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperRequest) GoString() string {
	return s.String()
}

func (s *GetDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
