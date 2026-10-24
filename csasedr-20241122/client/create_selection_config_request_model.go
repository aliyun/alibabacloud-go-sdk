// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSelectionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateSelectionConfigRequest struct {
}

func (s CreateSelectionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSelectionConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSelectionConfigRequest) Validate() error {
	return dara.Validate(s)
}
