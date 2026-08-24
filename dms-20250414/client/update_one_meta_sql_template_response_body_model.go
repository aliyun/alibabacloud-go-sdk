// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaSqlTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OneMetaSqlTemplateView) *UpdateOneMetaSqlTemplateResponseBody
	GetData() *OneMetaSqlTemplateView
	SetErrorCode(v string) *UpdateOneMetaSqlTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateOneMetaSqlTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateOneMetaSqlTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOneMetaSqlTemplateResponseBody
	GetSuccess() *bool
}

type UpdateOneMetaSqlTemplateResponseBody struct {
	Data         *OneMetaSqlTemplateView `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOneMetaSqlTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaSqlTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaSqlTemplateResponseBody) GetData() *OneMetaSqlTemplateView {
	return s.Data
}

func (s *UpdateOneMetaSqlTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateOneMetaSqlTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateOneMetaSqlTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOneMetaSqlTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOneMetaSqlTemplateResponseBody) SetData(v *OneMetaSqlTemplateView) *UpdateOneMetaSqlTemplateResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponseBody) SetErrorCode(v string) *UpdateOneMetaSqlTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponseBody) SetErrorMessage(v string) *UpdateOneMetaSqlTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponseBody) SetRequestId(v string) *UpdateOneMetaSqlTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponseBody) SetSuccess(v bool) *UpdateOneMetaSqlTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
