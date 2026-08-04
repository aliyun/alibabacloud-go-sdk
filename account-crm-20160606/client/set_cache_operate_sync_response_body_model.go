// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCacheOperateSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetCacheOperateSyncResponseBody
	GetCode() *string
	SetData(v string) *SetCacheOperateSyncResponseBody
	GetData() *string
	SetMessage(v string) *SetCacheOperateSyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetCacheOperateSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetCacheOperateSyncResponseBody
	GetSuccess() *bool
}

type SetCacheOperateSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCacheOperateSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCacheOperateSyncResponseBody) GoString() string {
	return s.String()
}

func (s *SetCacheOperateSyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetCacheOperateSyncResponseBody) GetData() *string {
	return s.Data
}

func (s *SetCacheOperateSyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetCacheOperateSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCacheOperateSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetCacheOperateSyncResponseBody) SetCode(v string) *SetCacheOperateSyncResponseBody {
	s.Code = &v
	return s
}

func (s *SetCacheOperateSyncResponseBody) SetData(v string) *SetCacheOperateSyncResponseBody {
	s.Data = &v
	return s
}

func (s *SetCacheOperateSyncResponseBody) SetMessage(v string) *SetCacheOperateSyncResponseBody {
	s.Message = &v
	return s
}

func (s *SetCacheOperateSyncResponseBody) SetRequestId(v string) *SetCacheOperateSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCacheOperateSyncResponseBody) SetSuccess(v bool) *SetCacheOperateSyncResponseBody {
	s.Success = &v
	return s
}

func (s *SetCacheOperateSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
