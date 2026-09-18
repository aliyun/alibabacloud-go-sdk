// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeScanRequest struct {
}

func (s DescribeScanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanRequest) Validate() error {
	return dara.Validate(s)
}
