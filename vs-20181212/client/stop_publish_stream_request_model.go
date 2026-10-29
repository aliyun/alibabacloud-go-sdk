// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPublishStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopPublishStreamRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *StopPublishStreamRequest
	GetOwnerId() *int64
}

type StopPublishStreamRequest struct {
	// The ID of the cloud application service instance for which to stop desktop capture and streaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StopPublishStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPublishStreamRequest) GoString() string {
	return s.String()
}

func (s *StopPublishStreamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopPublishStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopPublishStreamRequest) SetInstanceId(v string) *StopPublishStreamRequest {
	s.InstanceId = &v
	return s
}

func (s *StopPublishStreamRequest) SetOwnerId(v int64) *StopPublishStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopPublishStreamRequest) Validate() error {
	return dara.Validate(s)
}
