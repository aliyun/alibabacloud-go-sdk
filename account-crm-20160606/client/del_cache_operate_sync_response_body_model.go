// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelCacheOperateSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelCacheOperateSyncResponseBody
	GetCode() *string
	SetData(v string) *DelCacheOperateSyncResponseBody
	GetData() *string
	SetMessage(v string) *DelCacheOperateSyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelCacheOperateSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelCacheOperateSyncResponseBody
	GetSuccess() *bool
}

type DelCacheOperateSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelCacheOperateSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelCacheOperateSyncResponseBody) GoString() string {
	return s.String()
}

func (s *DelCacheOperateSyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelCacheOperateSyncResponseBody) GetData() *string {
	return s.Data
}

func (s *DelCacheOperateSyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelCacheOperateSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelCacheOperateSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelCacheOperateSyncResponseBody) SetCode(v string) *DelCacheOperateSyncResponseBody {
	s.Code = &v
	return s
}

func (s *DelCacheOperateSyncResponseBody) SetData(v string) *DelCacheOperateSyncResponseBody {
	s.Data = &v
	return s
}

func (s *DelCacheOperateSyncResponseBody) SetMessage(v string) *DelCacheOperateSyncResponseBody {
	s.Message = &v
	return s
}

func (s *DelCacheOperateSyncResponseBody) SetRequestId(v string) *DelCacheOperateSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelCacheOperateSyncResponseBody) SetSuccess(v bool) *DelCacheOperateSyncResponseBody {
	s.Success = &v
	return s
}

func (s *DelCacheOperateSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
