// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSipTrunksRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListSipTrunksRequest struct {
}

func (s ListSipTrunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSipTrunksRequest) GoString() string {
	return s.String()
}

func (s *ListSipTrunksRequest) Validate() error {
	return dara.Validate(s)
}
