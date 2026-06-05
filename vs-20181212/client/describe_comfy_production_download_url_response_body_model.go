// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyProductionDownloadUrlResponseBody
	GetCode() *int64
	SetDownloadUrl(v string) *DescribeComfyProductionDownloadUrlResponseBody
	GetDownloadUrl() *string
	SetExpiredTime(v string) *DescribeComfyProductionDownloadUrlResponseBody
	GetExpiredTime() *string
	SetMessage(v string) *DescribeComfyProductionDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeComfyProductionDownloadUrlResponseBody
	GetRequestId() *string
}

type DescribeComfyProductionDownloadUrlResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://xxx.xxx.xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 2029-03-28T16:00Z
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

func (s DescribeComfyProductionDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) SetCode(v int64) *DescribeComfyProductionDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) SetDownloadUrl(v string) *DescribeComfyProductionDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) SetExpiredTime(v string) *DescribeComfyProductionDownloadUrlResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) SetMessage(v string) *DescribeComfyProductionDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) SetRequestId(v string) *DescribeComfyProductionDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
