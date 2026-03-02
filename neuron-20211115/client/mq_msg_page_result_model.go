// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqMsgPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetMqMsgDigests(v []*MqMsgDigest) *MqMsgPageResult
	GetMqMsgDigests() []*MqMsgDigest
	SetIntTotal(v int32) *MqMsgPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MqMsgPageResult
	GetRequestId() *string
	SetTotal(v int64) *MqMsgPageResult
	GetTotal() *int64
}

type MqMsgPageResult struct {
	MqMsgDigests []*MqMsgDigest `json:"MqMsgDigests,omitempty" xml:"MqMsgDigests,omitempty" type:"Repeated"`
	IntTotal     *int32         `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId    *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MqMsgPageResult) String() string {
	return dara.Prettify(s)
}

func (s MqMsgPageResult) GoString() string {
	return s.String()
}

func (s *MqMsgPageResult) GetMqMsgDigests() []*MqMsgDigest {
	return s.MqMsgDigests
}

func (s *MqMsgPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MqMsgPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MqMsgPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MqMsgPageResult) SetMqMsgDigests(v []*MqMsgDigest) *MqMsgPageResult {
	s.MqMsgDigests = v
	return s
}

func (s *MqMsgPageResult) SetIntTotal(v int32) *MqMsgPageResult {
	s.IntTotal = &v
	return s
}

func (s *MqMsgPageResult) SetRequestId(v string) *MqMsgPageResult {
	s.RequestId = &v
	return s
}

func (s *MqMsgPageResult) SetTotal(v int64) *MqMsgPageResult {
	s.Total = &v
	return s
}

func (s *MqMsgPageResult) Validate() error {
	if s.MqMsgDigests != nil {
		for _, item := range s.MqMsgDigests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
