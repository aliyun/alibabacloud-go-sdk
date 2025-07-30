// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningStrategyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWarningStrategyConfigResponseBody
	GetCode() *string
	SetData(v string) *DeleteWarningStrategyConfigResponseBody
	GetData() *string
	SetMessage(v string) *DeleteWarningStrategyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWarningStrategyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWarningStrategyConfigResponseBody
	GetSuccess() *bool
}

type DeleteWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningStrategyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWarningStrategyConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteWarningStrategyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWarningStrategyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWarningStrategyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWarningStrategyConfigResponseBody) SetCode(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetData(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetMessage(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetRequestId(v string) *DeleteWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetSuccess(v bool) *DeleteWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
