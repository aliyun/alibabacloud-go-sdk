// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadFileResponseBody
	GetRequestId() *string
	SetUploadId(v int64) *UploadFileResponseBody
	GetUploadId() *int64
}

type UploadFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadId  *int64  `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s UploadFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadFileResponseBody) GetUploadId() *int64 {
	return s.UploadId
}

func (s *UploadFileResponseBody) SetRequestId(v string) *UploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFileResponseBody) SetUploadId(v int64) *UploadFileResponseBody {
	s.UploadId = &v
	return s
}

func (s *UploadFileResponseBody) Validate() error {
	return dara.Validate(s)
}
