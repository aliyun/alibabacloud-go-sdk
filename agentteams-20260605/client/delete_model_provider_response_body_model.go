// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelProviderResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteModelProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelProviderResponseBody
	GetSuccess() *bool
}

type DeleteModelProviderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelProviderResponseBody) SetCode(v string) *DeleteModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelProviderResponseBody) SetMessage(v string) *DeleteModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelProviderResponseBody) SetRequestId(v string) *DeleteModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelProviderResponseBody) SetSuccess(v bool) *DeleteModelProviderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
