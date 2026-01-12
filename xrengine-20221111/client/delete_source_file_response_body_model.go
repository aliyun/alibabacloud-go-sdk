// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSourceFileResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSourceFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSourceFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSourceFileResponseBody
	GetSuccess() *bool
}

type DeleteSourceFileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSourceFileResponseBody) SetCode(v string) *DeleteSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSourceFileResponseBody) SetMessage(v string) *DeleteSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSourceFileResponseBody) SetRequestId(v string) *DeleteSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceFileResponseBody) SetSuccess(v bool) *DeleteSourceFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
