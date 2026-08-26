// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentThemeResponseBodyData) *CreateDataAgentThemeResponseBody
	GetData() *CreateDataAgentThemeResponseBodyData
	SetErrorCode(v string) *CreateDataAgentThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentThemeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentThemeResponseBody
	GetSuccess() *bool
}

type CreateDataAgentThemeResponseBody struct {
	// The response struct.
	Data *CreateDataAgentThemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request is abnormal.
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
	// The request ID, which is used to locate logs and troubleshoot issues.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataAgentThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentThemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentThemeResponseBody) GetData() *CreateDataAgentThemeResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentThemeResponseBody) SetData(v *CreateDataAgentThemeResponseBodyData) *CreateDataAgentThemeResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentThemeResponseBody) SetErrorCode(v string) *CreateDataAgentThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentThemeResponseBody) SetErrorMessage(v string) *CreateDataAgentThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentThemeResponseBody) SetRequestId(v string) *CreateDataAgentThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentThemeResponseBody) SetSuccess(v bool) *CreateDataAgentThemeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentThemeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentThemeResponseBodyData struct {
	// The OSS key of the theme file (dart/{uid}/{theme_id}/theme.zip, verified to exist before being stored in the database).
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The business identifier of the theme.
	//
	// example:
	//
	// 0f8b2c1d************9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s CreateDataAgentThemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentThemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentThemeResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateDataAgentThemeResponseBodyData) GetThemeId() *string {
	return s.ThemeId
}

func (s *CreateDataAgentThemeResponseBodyData) SetFilePath(v string) *CreateDataAgentThemeResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *CreateDataAgentThemeResponseBodyData) SetThemeId(v string) *CreateDataAgentThemeResponseBodyData {
	s.ThemeId = &v
	return s
}

func (s *CreateDataAgentThemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
