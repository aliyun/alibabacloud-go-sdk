// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSqlLogConfigRequest
	GetInstanceId() *string
}

type DescribeSqlLogConfigRequest struct {
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSqlLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogConfigRequest) SetInstanceId(v string) *DescribeSqlLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
