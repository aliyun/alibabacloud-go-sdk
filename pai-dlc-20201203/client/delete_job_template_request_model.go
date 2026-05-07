// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteJobTemplateRequest struct {
}

func (s DeleteJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
