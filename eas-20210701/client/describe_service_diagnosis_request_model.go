// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceDiagnosisRequest struct {
}

func (s DescribeServiceDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
