// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNoticeInstanceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *NoticeInstanceUserRequest
	GetInstanceId() *int64
	SetNoticeParam(v string) *NoticeInstanceUserRequest
	GetNoticeParam() *string
	SetNoticeType(v int64) *NoticeInstanceUserRequest
	GetNoticeType() *int64
}

type NoticeInstanceUserRequest struct {
	// example:
	//
	// 5000000264872
	InstanceId  *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NoticeParam *string `json:"NoticeParam,omitempty" xml:"NoticeParam,omitempty"`
	// example:
	//
	// 1
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
}

func (s NoticeInstanceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s NoticeInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *NoticeInstanceUserRequest) GetNoticeParam() *string {
	return s.NoticeParam
}

func (s *NoticeInstanceUserRequest) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *NoticeInstanceUserRequest) SetInstanceId(v int64) *NoticeInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *NoticeInstanceUserRequest) SetNoticeParam(v string) *NoticeInstanceUserRequest {
	s.NoticeParam = &v
	return s
}

func (s *NoticeInstanceUserRequest) SetNoticeType(v int64) *NoticeInstanceUserRequest {
	s.NoticeType = &v
	return s
}

func (s *NoticeInstanceUserRequest) Validate() error {
	return dara.Validate(s)
}
