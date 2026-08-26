// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOneMetaOssieModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OssieModelView) *GetOneMetaOssieModelResponseBody
	GetData() *OssieModelView
	SetErrorCode(v string) *GetOneMetaOssieModelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOneMetaOssieModelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetOneMetaOssieModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOneMetaOssieModelResponseBody
	GetSuccess() *bool
}

type GetOneMetaOssieModelResponseBody struct {
	// The response struct.
	Data *OssieModelView `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
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

func (s GetOneMetaOssieModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOneMetaOssieModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetOneMetaOssieModelResponseBody) GetData() *OssieModelView {
	return s.Data
}

func (s *GetOneMetaOssieModelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOneMetaOssieModelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOneMetaOssieModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOneMetaOssieModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOneMetaOssieModelResponseBody) SetData(v *OssieModelView) *GetOneMetaOssieModelResponseBody {
	s.Data = v
	return s
}

func (s *GetOneMetaOssieModelResponseBody) SetErrorCode(v string) *GetOneMetaOssieModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOneMetaOssieModelResponseBody) SetErrorMessage(v string) *GetOneMetaOssieModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOneMetaOssieModelResponseBody) SetRequestId(v string) *GetOneMetaOssieModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOneMetaOssieModelResponseBody) SetSuccess(v bool) *GetOneMetaOssieModelResponseBody {
	s.Success = &v
	return s
}

func (s *GetOneMetaOssieModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
