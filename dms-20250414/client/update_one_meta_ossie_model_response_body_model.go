// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaOssieModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OssieModelView) *UpdateOneMetaOssieModelResponseBody
	GetData() *OssieModelView
	SetErrorCode(v string) *UpdateOneMetaOssieModelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateOneMetaOssieModelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateOneMetaOssieModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOneMetaOssieModelResponseBody
	GetSuccess() *bool
}

type UpdateOneMetaOssieModelResponseBody struct {
	// The response struct.
	Data *OssieModelView `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOneMetaOssieModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaOssieModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaOssieModelResponseBody) GetData() *OssieModelView {
	return s.Data
}

func (s *UpdateOneMetaOssieModelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateOneMetaOssieModelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateOneMetaOssieModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOneMetaOssieModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOneMetaOssieModelResponseBody) SetData(v *OssieModelView) *UpdateOneMetaOssieModelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOneMetaOssieModelResponseBody) SetErrorCode(v string) *UpdateOneMetaOssieModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateOneMetaOssieModelResponseBody) SetErrorMessage(v string) *UpdateOneMetaOssieModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateOneMetaOssieModelResponseBody) SetRequestId(v string) *UpdateOneMetaOssieModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOneMetaOssieModelResponseBody) SetSuccess(v bool) *UpdateOneMetaOssieModelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOneMetaOssieModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
