// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSearchParametersRequest
	GetDBClusterId() *string
}

type DescribeSearchParametersRequest struct {
	// The ID of the PolarDB primary cluster, which is used to locate the Search cluster associated with the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeSearchParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSearchParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSearchParametersRequest) SetDBClusterId(v string) *DescribeSearchParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSearchParametersRequest) Validate() error {
	return dara.Validate(s)
}
