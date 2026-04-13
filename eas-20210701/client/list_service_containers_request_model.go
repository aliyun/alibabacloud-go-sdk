// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceContainersRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListServiceContainersRequest struct {
}

func (s ListServiceContainersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceContainersRequest) GoString() string {
	return s.String()
}

func (s *ListServiceContainersRequest) Validate() error {
	return dara.Validate(s)
}
