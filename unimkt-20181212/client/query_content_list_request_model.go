// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandUserId(v int64) *QueryContentListRequest
	GetBrandUserId() *int64
	SetBrandUserNick(v string) *QueryContentListRequest
	GetBrandUserNick() *string
	SetChannelId(v string) *QueryContentListRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *QueryContentListRequest
	GetProxyUserId() *int64
	SetTaskBizType(v string) *QueryContentListRequest
	GetTaskBizType() *string
	SetTaskType(v string) *QueryContentListRequest
	GetTaskType() *string
}

type QueryContentListRequest struct {
	BrandUserId   *int64  `json:"BrandUserId,omitempty" xml:"BrandUserId,omitempty"`
	BrandUserNick *string `json:"BrandUserNick,omitempty" xml:"BrandUserNick,omitempty"`
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// This parameter is required.
	TaskBizType *string `json:"TaskBizType,omitempty" xml:"TaskBizType,omitempty"`
	// This parameter is required.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryContentListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContentListRequest) GoString() string {
	return s.String()
}

func (s *QueryContentListRequest) GetBrandUserId() *int64 {
	return s.BrandUserId
}

func (s *QueryContentListRequest) GetBrandUserNick() *string {
	return s.BrandUserNick
}

func (s *QueryContentListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryContentListRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *QueryContentListRequest) GetTaskBizType() *string {
	return s.TaskBizType
}

func (s *QueryContentListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryContentListRequest) SetBrandUserId(v int64) *QueryContentListRequest {
	s.BrandUserId = &v
	return s
}

func (s *QueryContentListRequest) SetBrandUserNick(v string) *QueryContentListRequest {
	s.BrandUserNick = &v
	return s
}

func (s *QueryContentListRequest) SetChannelId(v string) *QueryContentListRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryContentListRequest) SetProxyUserId(v int64) *QueryContentListRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryContentListRequest) SetTaskBizType(v string) *QueryContentListRequest {
	s.TaskBizType = &v
	return s
}

func (s *QueryContentListRequest) SetTaskType(v string) *QueryContentListRequest {
	s.TaskType = &v
	return s
}

func (s *QueryContentListRequest) Validate() error {
	return dara.Validate(s)
}
