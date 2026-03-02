// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Folder) *UpdateFolderResponseBody
	GetData() *Folder
	SetErrorCode(v string) *UpdateFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFolderResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateFolderResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFolderResponseBody
	GetSuccess() *bool
}

type UpdateFolderResponseBody struct {
	// The information returned after the folder is updated.
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) GetData() *Folder {
	return s.Data
}

func (s *UpdateFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFolderResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFolderResponseBody) SetData(v *Folder) *UpdateFolderResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFolderResponseBody) SetErrorCode(v string) *UpdateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetErrorMessage(v string) *UpdateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFolderResponseBody) SetHttpCode(v int32) *UpdateFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFolderResponseBody) SetSuccess(v bool) *UpdateFolderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFolderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
