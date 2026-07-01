// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Folder) *GetFolderResponseBody
	GetData() *Folder
	SetErrorCode(v string) *GetFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFolderResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetFolderResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFolderResponseBody
	GetSuccess() *bool
}

type GetFolderResponseBody struct {
	// Folder data structure.
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// - If success is false, the business error code is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// - If success is false, the business error message is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Indicates task success or failure. If the parameter is not empty and not 200, it means failure.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Request ID.
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

func (s GetFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) GetData() *Folder {
	return s.Data
}

func (s *GetFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFolderResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFolderResponseBody) SetData(v *Folder) *GetFolderResponseBody {
	s.Data = v
	return s
}

func (s *GetFolderResponseBody) SetErrorCode(v string) *GetFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFolderResponseBody) SetErrorMessage(v string) *GetFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFolderResponseBody) SetHttpCode(v int32) *GetFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFolderResponseBody) SetSuccess(v bool) *GetFolderResponseBody {
	s.Success = &v
	return s
}

func (s *GetFolderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
