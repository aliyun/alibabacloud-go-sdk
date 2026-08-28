// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMseNacosSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListMseNacosSourcesRequest struct {
}

func (s ListMseNacosSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMseNacosSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMseNacosSourcesRequest) Validate() error {
	return dara.Validate(s)
}
