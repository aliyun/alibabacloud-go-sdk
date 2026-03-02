// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqTopicSubsPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetMqTopicSubsDigests(v []*MqTopicSubsDigest) *MqTopicSubsPageResult
	GetMqTopicSubsDigests() []*MqTopicSubsDigest
	SetIntTotal(v int32) *MqTopicSubsPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MqTopicSubsPageResult
	GetRequestId() *string
	SetTotal(v int64) *MqTopicSubsPageResult
	GetTotal() *int64
}

type MqTopicSubsPageResult struct {
	MqTopicSubsDigests []*MqTopicSubsDigest `json:"MqTopicSubsDigests,omitempty" xml:"MqTopicSubsDigests,omitempty" type:"Repeated"`
	IntTotal           *int32               `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId          *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MqTopicSubsPageResult) String() string {
	return dara.Prettify(s)
}

func (s MqTopicSubsPageResult) GoString() string {
	return s.String()
}

func (s *MqTopicSubsPageResult) GetMqTopicSubsDigests() []*MqTopicSubsDigest {
	return s.MqTopicSubsDigests
}

func (s *MqTopicSubsPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MqTopicSubsPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MqTopicSubsPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MqTopicSubsPageResult) SetMqTopicSubsDigests(v []*MqTopicSubsDigest) *MqTopicSubsPageResult {
	s.MqTopicSubsDigests = v
	return s
}

func (s *MqTopicSubsPageResult) SetIntTotal(v int32) *MqTopicSubsPageResult {
	s.IntTotal = &v
	return s
}

func (s *MqTopicSubsPageResult) SetRequestId(v string) *MqTopicSubsPageResult {
	s.RequestId = &v
	return s
}

func (s *MqTopicSubsPageResult) SetTotal(v int64) *MqTopicSubsPageResult {
	s.Total = &v
	return s
}

func (s *MqTopicSubsPageResult) Validate() error {
	if s.MqTopicSubsDigests != nil {
		for _, item := range s.MqTopicSubsDigests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
