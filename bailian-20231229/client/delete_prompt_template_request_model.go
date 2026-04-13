// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePromptTemplateRequest struct {
}

func (s DeletePromptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeletePromptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
