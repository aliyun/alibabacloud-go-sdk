// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqMsgDigest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MqMsgDigest
	GetId() *int64
	SetRetryCount(v int64) *MqMsgDigest
	GetRetryCount() *int64
	SetStoreSize(v int32) *MqMsgDigest
	GetStoreSize() *int32
	SetStoreTime(v string) *MqMsgDigest
	GetStoreTime() *string
	SetTags(v string) *MqMsgDigest
	GetTags() *string
}

type MqMsgDigest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RetryCount *int64  `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	StoreSize  *int32  `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTime  *string `json:"StoreTime,omitempty" xml:"StoreTime,omitempty"`
	Tags       *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s MqMsgDigest) String() string {
	return dara.Prettify(s)
}

func (s MqMsgDigest) GoString() string {
	return s.String()
}

func (s *MqMsgDigest) GetId() *int64 {
	return s.Id
}

func (s *MqMsgDigest) GetRetryCount() *int64 {
	return s.RetryCount
}

func (s *MqMsgDigest) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *MqMsgDigest) GetStoreTime() *string {
	return s.StoreTime
}

func (s *MqMsgDigest) GetTags() *string {
	return s.Tags
}

func (s *MqMsgDigest) SetId(v int64) *MqMsgDigest {
	s.Id = &v
	return s
}

func (s *MqMsgDigest) SetRetryCount(v int64) *MqMsgDigest {
	s.RetryCount = &v
	return s
}

func (s *MqMsgDigest) SetStoreSize(v int32) *MqMsgDigest {
	s.StoreSize = &v
	return s
}

func (s *MqMsgDigest) SetStoreTime(v string) *MqMsgDigest {
	s.StoreTime = &v
	return s
}

func (s *MqMsgDigest) SetTags(v string) *MqMsgDigest {
	s.Tags = &v
	return s
}

func (s *MqMsgDigest) Validate() error {
	return dara.Validate(s)
}
