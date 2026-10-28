// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReadonlyModeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetReadonlyModeRequest struct {
}

func (s GetReadonlyModeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReadonlyModeRequest) GoString() string {
	return s.String()
}

func (s *GetReadonlyModeRequest) Validate() error {
	return dara.Validate(s)
}
