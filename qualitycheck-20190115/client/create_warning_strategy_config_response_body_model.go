// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningStrategyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWarningStrategyConfigResponseBody
	GetCode() *string
	SetData(v int64) *CreateWarningStrategyConfigResponseBody
	GetData() *int64
	SetMessage(v string) *CreateWarningStrategyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWarningStrategyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWarningStrategyConfigResponseBody
	GetSuccess() *bool
}

type CreateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningStrategyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWarningStrategyConfigResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateWarningStrategyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWarningStrategyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarningStrategyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWarningStrategyConfigResponseBody) SetCode(v string) *CreateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetData(v int64) *CreateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetMessage(v string) *CreateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetRequestId(v string) *CreateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetSuccess(v bool) *CreateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
