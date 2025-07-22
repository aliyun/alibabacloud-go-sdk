// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoLiveStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAutoLiveStreamRuleRequest
	GetAppId() *string
	SetOwnerId(v int64) *DescribeAutoLiveStreamRuleRequest
	GetOwnerId() *int64
}

type DescribeAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAutoLiveStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAutoLiveStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoLiveStreamRuleRequest) SetAppId(v string) *DescribeAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DescribeAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
