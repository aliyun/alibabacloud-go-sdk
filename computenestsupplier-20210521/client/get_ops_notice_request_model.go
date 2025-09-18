// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoticeId(v string) *GetOpsNoticeRequest
	GetNoticeId() *string
	SetRegionId(v string) *GetOpsNoticeRequest
	GetRegionId() *string
}

type GetOpsNoticeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// notice-2338dxxxxxx
	NoticeId *string `json:"NoticeId,omitempty" xml:"NoticeId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOpsNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpsNoticeRequest) GoString() string {
	return s.String()
}

func (s *GetOpsNoticeRequest) GetNoticeId() *string {
	return s.NoticeId
}

func (s *GetOpsNoticeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOpsNoticeRequest) SetNoticeId(v string) *GetOpsNoticeRequest {
	s.NoticeId = &v
	return s
}

func (s *GetOpsNoticeRequest) SetRegionId(v string) *GetOpsNoticeRequest {
	s.RegionId = &v
	return s
}

func (s *GetOpsNoticeRequest) Validate() error {
	return dara.Validate(s)
}
