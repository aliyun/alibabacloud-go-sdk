// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTTSConfigRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *DescribeTTSConfigRequest
	GetInstanceOwnerId() *int64
}

type DescribeTTSConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dc437bba-5a25-4bbc-b4c2-f268864bebb5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1971226538081821
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s DescribeTTSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTTSConfigRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *DescribeTTSConfigRequest) SetInstanceId(v string) *DescribeTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTTSConfigRequest) SetInstanceOwnerId(v int64) *DescribeTTSConfigRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DescribeTTSConfigRequest) Validate() error {
	return dara.Validate(s)
}
