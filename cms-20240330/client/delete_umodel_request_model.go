// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteUmodelRequest struct {
}

func (s DeleteUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelRequest) GoString() string {
	return s.String()
}

func (s *DeleteUmodelRequest) Validate() error {
	return dara.Validate(s)
}
