// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMeteringDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetering(v string) *PushMeteringDataRequest
	GetMetering() *string
}

type PushMeteringDataRequest struct {
	// example:
	//
	// [{"InstanceId":"1000001","StartTime":"100000000","EndTime":"100000010","Entities":[{"Key":"PeriodMin","Value":"96","meteringAssit":"cmapi00060317-PeriodMin-4"}]}]
	Metering *string `json:"Metering,omitempty" xml:"Metering,omitempty"`
}

func (s PushMeteringDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) GetMetering() *string {
	return s.Metering
}

func (s *PushMeteringDataRequest) SetMetering(v string) *PushMeteringDataRequest {
	s.Metering = &v
	return s
}

func (s *PushMeteringDataRequest) Validate() error {
	return dara.Validate(s)
}
