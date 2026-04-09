// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTemplateRequest struct {
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
