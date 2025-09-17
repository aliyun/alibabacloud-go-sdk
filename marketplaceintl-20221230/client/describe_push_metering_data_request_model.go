// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePushMeteringDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPushOrderBizId(v string) *DescribePushMeteringDataRequest
	GetPushOrderBizId() *string
}

type DescribePushMeteringDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pushData123456
	PushOrderBizId *string `json:"PushOrderBizId,omitempty" xml:"PushOrderBizId,omitempty"`
}

func (s DescribePushMeteringDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *DescribePushMeteringDataRequest) GetPushOrderBizId() *string {
	return s.PushOrderBizId
}

func (s *DescribePushMeteringDataRequest) SetPushOrderBizId(v string) *DescribePushMeteringDataRequest {
	s.PushOrderBizId = &v
	return s
}

func (s *DescribePushMeteringDataRequest) Validate() error {
	return dara.Validate(s)
}
