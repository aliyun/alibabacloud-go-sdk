// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMeteringDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PushMeteringDataResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *PushMeteringDataResponseBody
	GetDynamicMessage() *string
	SetForceFatal(v bool) *PushMeteringDataResponseBody
	GetForceFatal() *bool
	SetMessage(v string) *PushMeteringDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushMeteringDataResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushMeteringDataResponseBody
	GetResult() *bool
	SetSuccess(v bool) *PushMeteringDataResponseBody
	GetSuccess() *bool
}

type PushMeteringDataResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ForceFatal     *bool   `json:"ForceFatal,omitempty" xml:"ForceFatal,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *PushMeteringDataResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PushMeteringDataResponseBody) GetForceFatal() *bool {
	return s.ForceFatal
}

func (s *PushMeteringDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushMeteringDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushMeteringDataResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushMeteringDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushMeteringDataResponseBody) SetCode(v string) *PushMeteringDataResponseBody {
	s.Code = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetDynamicMessage(v string) *PushMeteringDataResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetForceFatal(v bool) *PushMeteringDataResponseBody {
	s.ForceFatal = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetMessage(v string) *PushMeteringDataResponseBody {
	s.Message = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetResult(v bool) *PushMeteringDataResponseBody {
	s.Result = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetSuccess(v bool) *PushMeteringDataResponseBody {
	s.Success = &v
	return s
}

func (s *PushMeteringDataResponseBody) Validate() error {
	return dara.Validate(s)
}
