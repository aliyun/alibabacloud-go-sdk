// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQuotaTipRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetQuotaTipRequest
	GetRegionId() *string
}

type GetQuotaTipRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-i7m2wpm5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQuotaTipRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTipRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaTipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQuotaTipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaTipRequest) SetInstanceId(v string) *GetQuotaTipRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQuotaTipRequest) SetRegionId(v string) *GetQuotaTipRequest {
	s.RegionId = &v
	return s
}

func (s *GetQuotaTipRequest) Validate() error {
	return dara.Validate(s)
}
