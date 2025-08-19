// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceMNSParameters interface {
	dara.Model
	String() string
	GoString() string
	SetIsBase64Decode(v bool) *SourceMNSParameters
	GetIsBase64Decode() *bool
	SetQueueName(v string) *SourceMNSParameters
	GetQueueName() *string
	SetRegionId(v string) *SourceMNSParameters
	GetRegionId() *string
}

type SourceMNSParameters struct {
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceMNSParameters) GoString() string {
	return s.String()
}

func (s *SourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *SourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *SourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceMNSParameters) SetIsBase64Decode(v bool) *SourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *SourceMNSParameters) SetQueueName(v string) *SourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *SourceMNSParameters) SetRegionId(v string) *SourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *SourceMNSParameters) Validate() error {
	return dara.Validate(s)
}
