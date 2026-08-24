// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaSqlTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteOneMetaSqlTemplateResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOneMetaSqlTemplateResponseBody
	GetSuccess() *bool
}

type DeleteOneMetaSqlTemplateResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOneMetaSqlTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaSqlTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetData(v bool) *DeleteOneMetaSqlTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetErrorCode(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetErrorMessage(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetRequestId(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetSuccess(v bool) *DeleteOneMetaSqlTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
