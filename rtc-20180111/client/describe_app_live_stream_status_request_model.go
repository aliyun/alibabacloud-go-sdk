// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLiveStreamStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppLiveStreamStatusRequest
	GetAppId() *string
	SetClientToken(v string) *DescribeAppLiveStreamStatusRequest
	GetClientToken() *string
}

type DescribeAppLiveStreamStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeAppLiveStreamStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLiveStreamStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppLiveStreamStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppLiveStreamStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAppLiveStreamStatusRequest) SetAppId(v string) *DescribeAppLiveStreamStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppLiveStreamStatusRequest) SetClientToken(v string) *DescribeAppLiveStreamStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAppLiveStreamStatusRequest) Validate() error {
	return dara.Validate(s)
}
