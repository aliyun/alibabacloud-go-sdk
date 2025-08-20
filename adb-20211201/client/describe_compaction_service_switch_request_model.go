// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCompactionServiceSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeCompactionServiceSwitchRequest
	GetDBClusterId() *string
}

type DescribeCompactionServiceSwitchRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp149vz49b36t****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeCompactionServiceSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCompactionServiceSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeCompactionServiceSwitchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeCompactionServiceSwitchRequest) SetDBClusterId(v string) *DescribeCompactionServiceSwitchRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeCompactionServiceSwitchRequest) Validate() error {
	return dara.Validate(s)
}
