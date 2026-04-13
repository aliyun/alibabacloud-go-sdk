// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceFaultInjectionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListServiceInstanceFaultInjectionInfoRequest struct {
}

func (s ListServiceInstanceFaultInjectionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceFaultInjectionInfoRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceFaultInjectionInfoRequest) Validate() error {
	return dara.Validate(s)
}
