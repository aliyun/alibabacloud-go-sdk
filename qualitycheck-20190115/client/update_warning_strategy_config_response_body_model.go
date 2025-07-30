// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningStrategyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWarningStrategyConfigResponseBody
	GetCode() *string
	SetData(v string) *UpdateWarningStrategyConfigResponseBody
	GetData() *string
	SetMessage(v string) *UpdateWarningStrategyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWarningStrategyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWarningStrategyConfigResponseBody
	GetSuccess() *bool
}

type UpdateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningStrategyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWarningStrategyConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateWarningStrategyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWarningStrategyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWarningStrategyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWarningStrategyConfigResponseBody) SetCode(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetData(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetMessage(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetRequestId(v string) *UpdateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetSuccess(v bool) *UpdateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
