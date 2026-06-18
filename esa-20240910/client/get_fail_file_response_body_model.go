// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFailFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadLink(v string) *GetFailFileResponseBody
	GetDownloadLink() *string
	SetRequestId(v string) *GetFailFileResponseBody
	GetRequestId() *string
}

type GetFailFileResponseBody struct {
	// The download link for the failed task file.
	//
	// example:
	//
	// https://****.oss-cn-shenzhen.aliyuncs.com/fail-1593805857882113?Expires=1708659191&OSSAccessKeyId=**********&Signature=**********
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFailFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFailFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFailFileResponseBody) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *GetFailFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFailFileResponseBody) SetDownloadLink(v string) *GetFailFileResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *GetFailFileResponseBody) SetRequestId(v string) *GetFailFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFailFileResponseBody) Validate() error {
	return dara.Validate(s)
}
