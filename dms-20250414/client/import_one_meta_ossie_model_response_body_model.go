// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneMetaOssieModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OssieModelView) *ImportOneMetaOssieModelResponseBody
	GetData() *OssieModelView
	SetErrorCode(v string) *ImportOneMetaOssieModelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ImportOneMetaOssieModelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ImportOneMetaOssieModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportOneMetaOssieModelResponseBody
	GetSuccess() *bool
}

type ImportOneMetaOssieModelResponseBody struct {
	Data         *OssieModelView `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportOneMetaOssieModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportOneMetaOssieModelResponseBody) GoString() string {
	return s.String()
}

func (s *ImportOneMetaOssieModelResponseBody) GetData() *OssieModelView {
	return s.Data
}

func (s *ImportOneMetaOssieModelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ImportOneMetaOssieModelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ImportOneMetaOssieModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportOneMetaOssieModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportOneMetaOssieModelResponseBody) SetData(v *OssieModelView) *ImportOneMetaOssieModelResponseBody {
	s.Data = v
	return s
}

func (s *ImportOneMetaOssieModelResponseBody) SetErrorCode(v string) *ImportOneMetaOssieModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportOneMetaOssieModelResponseBody) SetErrorMessage(v string) *ImportOneMetaOssieModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ImportOneMetaOssieModelResponseBody) SetRequestId(v string) *ImportOneMetaOssieModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportOneMetaOssieModelResponseBody) SetSuccess(v bool) *ImportOneMetaOssieModelResponseBody {
	s.Success = &v
	return s
}

func (s *ImportOneMetaOssieModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
