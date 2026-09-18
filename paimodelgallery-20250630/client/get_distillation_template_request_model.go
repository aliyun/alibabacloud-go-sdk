// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDistillationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDistillationTemplateRequest struct {
}

func (s GetDistillationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDistillationTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDistillationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
