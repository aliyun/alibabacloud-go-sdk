// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguDownloadSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MiguSourceDownloadDTO) *ModelRouterMiguDownloadSourceResponseBody
	GetData() *MiguSourceDownloadDTO
	SetErrCode(v string) *ModelRouterMiguDownloadSourceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterMiguDownloadSourceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterMiguDownloadSourceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterMiguDownloadSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterMiguDownloadSourceResponseBody
	GetSuccess() *bool
}

type ModelRouterMiguDownloadSourceResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *MiguSourceDownloadDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code. This value is passed through as-is from the console and is not mapped to platform standard error codes. Valid values: ERROR_PARAMETERS: Missing or invalid parameters. ERROR_NOT_FOUND: The source file does not exist or does not belong to the current account. ERROR_SOURCE_TRANSFERRING: The source file is being transferred. ERROR_SOURCE_EXPIRED: The source file has expired. ERROR_SERVER_INTERNAL: Internal error.
	//
	// example:
	//
	// ERROR_NOT_FOUND
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The business status code. Valid values: 0: Success. 400/403/404/409/410/500: Business error. Refer to errCode for details.
	//
	// example:
	//
	// 0
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterMiguDownloadSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguDownloadSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetData() *MiguSourceDownloadDTO {
	return s.Data
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterMiguDownloadSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetData(v *MiguSourceDownloadDTO) *ModelRouterMiguDownloadSourceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetErrCode(v string) *ModelRouterMiguDownloadSourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetErrMessage(v string) *ModelRouterMiguDownloadSourceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetHttpStatusCode(v int32) *ModelRouterMiguDownloadSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetRequestId(v string) *ModelRouterMiguDownloadSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) SetSuccess(v bool) *ModelRouterMiguDownloadSourceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
