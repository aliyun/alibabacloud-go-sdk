// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaOssieModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteOneMetaOssieModelResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteOneMetaOssieModelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteOneMetaOssieModelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteOneMetaOssieModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOneMetaOssieModelResponseBody
	GetSuccess() *bool
}

type DeleteOneMetaOssieModelResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOneMetaOssieModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaOssieModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaOssieModelResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteOneMetaOssieModelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteOneMetaOssieModelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteOneMetaOssieModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOneMetaOssieModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOneMetaOssieModelResponseBody) SetData(v bool) *DeleteOneMetaOssieModelResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponseBody) SetErrorCode(v string) *DeleteOneMetaOssieModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponseBody) SetErrorMessage(v string) *DeleteOneMetaOssieModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponseBody) SetRequestId(v string) *DeleteOneMetaOssieModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponseBody) SetSuccess(v bool) *DeleteOneMetaOssieModelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponseBody) Validate() error {
	return dara.Validate(s)
}
