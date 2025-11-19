// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeImageDetailRequest
	GetApiKey() *string
	SetImageId(v string) *DescribeImageDetailRequest
	GetImageId() *string
	SetSessionId(v string) *DescribeImageDetailRequest
	GetSessionId() *string
	SetShareCode(v string) *DescribeImageDetailRequest
	GetShareCode() *string
}

type DescribeImageDetailRequest struct {
	// This parameter is required.
	ApiKey    *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ShareCode *string `json:"ShareCode,omitempty" xml:"ShareCode,omitempty"`
}

func (s DescribeImageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageDetailRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeImageDetailRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageDetailRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeImageDetailRequest) GetShareCode() *string {
	return s.ShareCode
}

func (s *DescribeImageDetailRequest) SetApiKey(v string) *DescribeImageDetailRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeImageDetailRequest) SetImageId(v string) *DescribeImageDetailRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageDetailRequest) SetSessionId(v string) *DescribeImageDetailRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeImageDetailRequest) SetShareCode(v string) *DescribeImageDetailRequest {
	s.ShareCode = &v
	return s
}

func (s *DescribeImageDetailRequest) Validate() error {
	return dara.Validate(s)
}
