// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiInfosRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListApiInfosRequest struct {
}

func (s ListApiInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiInfosRequest) GoString() string {
	return s.String()
}

func (s *ListApiInfosRequest) Validate() error {
	return dara.Validate(s)
}
