// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncrByCacheOperateSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IncrByCacheOperateSyncResponseBody
	GetCode() *string
	SetData(v string) *IncrByCacheOperateSyncResponseBody
	GetData() *string
	SetMessage(v string) *IncrByCacheOperateSyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *IncrByCacheOperateSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IncrByCacheOperateSyncResponseBody
	GetSuccess() *bool
}

type IncrByCacheOperateSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IncrByCacheOperateSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IncrByCacheOperateSyncResponseBody) GoString() string {
	return s.String()
}

func (s *IncrByCacheOperateSyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *IncrByCacheOperateSyncResponseBody) GetData() *string {
	return s.Data
}

func (s *IncrByCacheOperateSyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IncrByCacheOperateSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IncrByCacheOperateSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IncrByCacheOperateSyncResponseBody) SetCode(v string) *IncrByCacheOperateSyncResponseBody {
	s.Code = &v
	return s
}

func (s *IncrByCacheOperateSyncResponseBody) SetData(v string) *IncrByCacheOperateSyncResponseBody {
	s.Data = &v
	return s
}

func (s *IncrByCacheOperateSyncResponseBody) SetMessage(v string) *IncrByCacheOperateSyncResponseBody {
	s.Message = &v
	return s
}

func (s *IncrByCacheOperateSyncResponseBody) SetRequestId(v string) *IncrByCacheOperateSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncrByCacheOperateSyncResponseBody) SetSuccess(v bool) *IncrByCacheOperateSyncResponseBody {
	s.Success = &v
	return s
}

func (s *IncrByCacheOperateSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
