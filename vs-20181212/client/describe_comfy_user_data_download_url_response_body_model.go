// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyUserDataDownloadUrlResponseBody
	GetCode() *int64
	SetDownloadUrl(v string) *DescribeComfyUserDataDownloadUrlResponseBody
	GetDownloadUrl() *string
	SetExpiredTime(v string) *DescribeComfyUserDataDownloadUrlResponseBody
	GetExpiredTime() *string
	SetMessage(v string) *DescribeComfyUserDataDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeComfyUserDataDownloadUrlResponseBody
	GetRequestId() *string
}

type DescribeComfyUserDataDownloadUrlResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://xxx.xxx.xxx.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 1752805579553
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComfyUserDataDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) SetCode(v int64) *DescribeComfyUserDataDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) SetDownloadUrl(v string) *DescribeComfyUserDataDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) SetExpiredTime(v string) *DescribeComfyUserDataDownloadUrlResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) SetMessage(v string) *DescribeComfyUserDataDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) SetRequestId(v string) *DescribeComfyUserDataDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
