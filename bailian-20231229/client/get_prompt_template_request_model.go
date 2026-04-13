// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPromptTemplateRequest struct {
}

func (s GetPromptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetPromptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
