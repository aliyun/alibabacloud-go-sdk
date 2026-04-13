// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceInstanceDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceInstanceDiagnosisRequest struct {
}

func (s DescribeServiceInstanceDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
