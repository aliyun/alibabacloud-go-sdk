// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqTopicPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetMqTopicDigests(v []*MqTopicDigest) *MqTopicPageResult
	GetMqTopicDigests() []*MqTopicDigest
	SetIntTotal(v int32) *MqTopicPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MqTopicPageResult
	GetRequestId() *string
	SetTotal(v int64) *MqTopicPageResult
	GetTotal() *int64
}

type MqTopicPageResult struct {
	MqTopicDigests []*MqTopicDigest `json:"MqTopicDigests,omitempty" xml:"MqTopicDigests,omitempty" type:"Repeated"`
	IntTotal       *int32           `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId      *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MqTopicPageResult) String() string {
	return dara.Prettify(s)
}

func (s MqTopicPageResult) GoString() string {
	return s.String()
}

func (s *MqTopicPageResult) GetMqTopicDigests() []*MqTopicDigest {
	return s.MqTopicDigests
}

func (s *MqTopicPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MqTopicPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MqTopicPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MqTopicPageResult) SetMqTopicDigests(v []*MqTopicDigest) *MqTopicPageResult {
	s.MqTopicDigests = v
	return s
}

func (s *MqTopicPageResult) SetIntTotal(v int32) *MqTopicPageResult {
	s.IntTotal = &v
	return s
}

func (s *MqTopicPageResult) SetRequestId(v string) *MqTopicPageResult {
	s.RequestId = &v
	return s
}

func (s *MqTopicPageResult) SetTotal(v int64) *MqTopicPageResult {
	s.Total = &v
	return s
}

func (s *MqTopicPageResult) Validate() error {
	if s.MqTopicDigests != nil {
		for _, item := range s.MqTopicDigests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
