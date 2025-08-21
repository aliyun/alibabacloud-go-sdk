// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPublishStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartPublishStreamRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *StartPublishStreamRequest
	GetOwnerId() *int64
	SetPublishUrl(v string) *StartPublishStreamRequest
	GetPublishUrl() *string
}

type StartPublishStreamRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PublishUrl *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
}

func (s StartPublishStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPublishStreamRequest) GoString() string {
	return s.String()
}

func (s *StartPublishStreamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartPublishStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartPublishStreamRequest) GetPublishUrl() *string {
	return s.PublishUrl
}

func (s *StartPublishStreamRequest) SetInstanceId(v string) *StartPublishStreamRequest {
	s.InstanceId = &v
	return s
}

func (s *StartPublishStreamRequest) SetOwnerId(v int64) *StartPublishStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartPublishStreamRequest) SetPublishUrl(v string) *StartPublishStreamRequest {
	s.PublishUrl = &v
	return s
}

func (s *StartPublishStreamRequest) Validate() error {
	return dara.Validate(s)
}
