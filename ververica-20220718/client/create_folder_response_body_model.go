// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Folder) *CreateFolderResponseBody
	GetData() *Folder
	SetErrorCode(v string) *CreateFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFolderResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateFolderResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFolderResponseBody
	GetSuccess() *bool
}

type CreateFolderResponseBody struct {
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) GetData() *Folder {
	return s.Data
}

func (s *CreateFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFolderResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFolderResponseBody) SetData(v *Folder) *CreateFolderResponseBody {
	s.Data = v
	return s
}

func (s *CreateFolderResponseBody) SetErrorCode(v string) *CreateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetErrorMessage(v string) *CreateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFolderResponseBody) SetHttpCode(v int32) *CreateFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponseBody) SetSuccess(v bool) *CreateFolderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFolderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
