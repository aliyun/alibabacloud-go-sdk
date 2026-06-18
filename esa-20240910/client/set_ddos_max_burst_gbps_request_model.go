// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDdosMaxBurstGbpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetDdosMaxBurstGbpsRequest
	GetInstanceId() *string
	SetMaxBurstGbps(v string) *SetDdosMaxBurstGbpsRequest
	GetMaxBurstGbps() *string
}

type SetDdosMaxBurstGbpsRequest struct {
	// The ID of the instance. You can call the [ListUserRatePlanInstances](https://help.aliyun.com/document_detail/2852398.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-ads11w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum burstable protection bandwidth for the DDoS instance in mainland China. The unit is Gbps.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	MaxBurstGbps *string `json:"MaxBurstGbps,omitempty" xml:"MaxBurstGbps,omitempty"`
}

func (s SetDdosMaxBurstGbpsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDdosMaxBurstGbpsRequest) GoString() string {
	return s.String()
}

func (s *SetDdosMaxBurstGbpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetDdosMaxBurstGbpsRequest) GetMaxBurstGbps() *string {
	return s.MaxBurstGbps
}

func (s *SetDdosMaxBurstGbpsRequest) SetInstanceId(v string) *SetDdosMaxBurstGbpsRequest {
	s.InstanceId = &v
	return s
}

func (s *SetDdosMaxBurstGbpsRequest) SetMaxBurstGbps(v string) *SetDdosMaxBurstGbpsRequest {
	s.MaxBurstGbps = &v
	return s
}

func (s *SetDdosMaxBurstGbpsRequest) Validate() error {
	return dara.Validate(s)
}
