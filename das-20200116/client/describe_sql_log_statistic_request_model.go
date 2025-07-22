// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSqlLogStatisticRequest
	GetInstanceId() *string
}

type DescribeSqlLogStatisticRequest struct {
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSqlLogStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogStatisticRequest) SetInstanceId(v string) *DescribeSqlLogStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogStatisticRequest) Validate() error {
	return dara.Validate(s)
}
