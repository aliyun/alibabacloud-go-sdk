// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BatchDeleteTopicsRequest
	GetInstanceId() *string
	SetRegionId(v string) *BatchDeleteTopicsRequest
	GetRegionId() *string
	SetTopics(v []*string) *BatchDeleteTopicsRequest
	GetTopics() []*string
}

type BatchDeleteTopicsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topics   []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s BatchDeleteTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchDeleteTopicsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchDeleteTopicsRequest) GetTopics() []*string {
	return s.Topics
}

func (s *BatchDeleteTopicsRequest) SetInstanceId(v string) *BatchDeleteTopicsRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchDeleteTopicsRequest) SetRegionId(v string) *BatchDeleteTopicsRequest {
	s.RegionId = &v
	return s
}

func (s *BatchDeleteTopicsRequest) SetTopics(v []*string) *BatchDeleteTopicsRequest {
	s.Topics = v
	return s
}

func (s *BatchDeleteTopicsRequest) Validate() error {
	return dara.Validate(s)
}
