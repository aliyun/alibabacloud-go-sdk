// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteTemplateRequest struct {
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
