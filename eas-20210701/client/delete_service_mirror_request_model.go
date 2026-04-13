// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMirrorRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteServiceMirrorRequest struct {
}

func (s DeleteServiceMirrorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorRequest) Validate() error {
	return dara.Validate(s)
}
