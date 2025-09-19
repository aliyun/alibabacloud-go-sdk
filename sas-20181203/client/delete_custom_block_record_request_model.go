// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomBlockRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCustomBlockRecordRequest
	GetId() *int64
	SetResourceOwnerId(v int64) *DeleteCustomBlockRecordRequest
	GetResourceOwnerId() *int64
}

type DeleteCustomBlockRecordRequest struct {
	// The ID of the IP address blocking policy.
	//
	// > You can call the [DescribeCustomBlockRecords](~~DescribeCustomBlockRecords~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 381**
	Id              *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCustomBlockRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomBlockRecordRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCustomBlockRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomBlockRecordRequest) SetId(v int64) *DeleteCustomBlockRecordRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomBlockRecordRequest) SetResourceOwnerId(v int64) *DeleteCustomBlockRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomBlockRecordRequest) Validate() error {
	return dara.Validate(s)
}
