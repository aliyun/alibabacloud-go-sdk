// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *GetConsumerListRequest
	GetConsumerId() *string
	SetCurrentPage(v int32) *GetConsumerListRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *GetConsumerListRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetConsumerListRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetConsumerListRequest
	GetRegionId() *string
}

type GetConsumerListRequest struct {
	// The name of the consumer group. If you do not configure this parameter, all consumer groups are queried.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h18sav****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to be returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerListRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetConsumerListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetConsumerListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetConsumerListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerListRequest) SetConsumerId(v string) *GetConsumerListRequest {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerListRequest) SetCurrentPage(v int32) *GetConsumerListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetConsumerListRequest) SetInstanceId(v string) *GetConsumerListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListRequest) SetPageSize(v int32) *GetConsumerListRequest {
	s.PageSize = &v
	return s
}

func (s *GetConsumerListRequest) SetRegionId(v string) *GetConsumerListRequest {
	s.RegionId = &v
	return s
}

func (s *GetConsumerListRequest) Validate() error {
	return dara.Validate(s)
}
