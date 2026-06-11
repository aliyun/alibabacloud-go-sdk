// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCmsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type OpenCmsServiceRequest struct {
}

func (s OpenCmsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenCmsServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCmsServiceRequest) Validate() error {
	return dara.Validate(s)
}
