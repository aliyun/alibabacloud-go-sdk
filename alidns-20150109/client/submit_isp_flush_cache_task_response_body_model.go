// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIspFlushCacheTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitIspFlushCacheTaskResponseBody
	GetRequestId() *string
}

type SubmitIspFlushCacheTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitIspFlushCacheTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIspFlushCacheTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIspFlushCacheTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIspFlushCacheTaskResponseBody) SetRequestId(v string) *SubmitIspFlushCacheTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIspFlushCacheTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
