// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *DescribeChatMessageRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeChatMessageRequest
	GetRegionId() *string
	SetSessionId(v string) *DescribeChatMessageRequest
	GetSessionId() *string
	SetTimezone(v string) *DescribeChatMessageRequest
	GetTimezone() *string
}

type DescribeChatMessageRequest struct {
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// df94eec5-3d95-435c-87d4-43e8bb3f9519
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s DescribeChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeChatMessageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeChatMessageRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeChatMessageRequest) SetQuery(v string) *DescribeChatMessageRequest {
	s.Query = &v
	return s
}

func (s *DescribeChatMessageRequest) SetRegionId(v string) *DescribeChatMessageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeChatMessageRequest) SetSessionId(v string) *DescribeChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeChatMessageRequest) SetTimezone(v string) *DescribeChatMessageRequest {
	s.Timezone = &v
	return s
}

func (s *DescribeChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
