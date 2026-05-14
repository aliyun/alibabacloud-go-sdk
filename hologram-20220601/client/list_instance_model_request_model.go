// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListInstanceModelRequest struct {
}

func (s ListInstanceModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceModelRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceModelRequest) Validate() error {
	return dara.Validate(s)
}
