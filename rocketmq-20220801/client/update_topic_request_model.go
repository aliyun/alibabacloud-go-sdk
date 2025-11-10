// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiteTopicExpiration(v int64) *UpdateTopicRequest
	GetLiteTopicExpiration() *int64
	SetMaxSendTps(v int64) *UpdateTopicRequest
	GetMaxSendTps() *int64
	SetRemark(v string) *UpdateTopicRequest
	GetRemark() *string
}

type UpdateTopicRequest struct {
	// example:
	//
	// 20
	LiteTopicExpiration *int64 `json:"liteTopicExpiration,omitempty" xml:"liteTopicExpiration,omitempty"`
	// Maximum send message tps
	//
	// example:
	//
	// 500
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// Updated remarks for the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequest) GetLiteTopicExpiration() *int64 {
	return s.LiteTopicExpiration
}

func (s *UpdateTopicRequest) GetMaxSendTps() *int64 {
	return s.MaxSendTps
}

func (s *UpdateTopicRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateTopicRequest) SetLiteTopicExpiration(v int64) *UpdateTopicRequest {
	s.LiteTopicExpiration = &v
	return s
}

func (s *UpdateTopicRequest) SetMaxSendTps(v int64) *UpdateTopicRequest {
	s.MaxSendTps = &v
	return s
}

func (s *UpdateTopicRequest) SetRemark(v string) *UpdateTopicRequest {
	s.Remark = &v
	return s
}

func (s *UpdateTopicRequest) Validate() error {
	return dara.Validate(s)
}
