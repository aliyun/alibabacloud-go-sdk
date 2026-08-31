// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguUploadSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MiguSourceUploadDTO) *ModelRouterMiguUploadSourceResponseBody
	GetData() *MiguSourceUploadDTO
	SetErrCode(v string) *ModelRouterMiguUploadSourceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterMiguUploadSourceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterMiguUploadSourceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterMiguUploadSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterMiguUploadSourceResponseBody
	GetSuccess() *bool
}

type ModelRouterMiguUploadSourceResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *MiguSourceUploadDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code. The console passes through the original value without mapping to standard platform error codes. Valid values:
	//
	// - ERROR_PARAMETERS: Missing or invalid parameters.
	//
	// - ERROR_SERVER_INTERNAL: Internal error.
	//
	// example:
	//
	// ERROR_PARAMETERS
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The business status code. Valid values:
	//
	// - 0: Success.
	//
	// - 400/403/500: Business error. For details, see errCode.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterMiguUploadSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguUploadSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetData() *MiguSourceUploadDTO {
	return s.Data
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterMiguUploadSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetData(v *MiguSourceUploadDTO) *ModelRouterMiguUploadSourceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetErrCode(v string) *ModelRouterMiguUploadSourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetErrMessage(v string) *ModelRouterMiguUploadSourceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetHttpStatusCode(v int32) *ModelRouterMiguUploadSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetRequestId(v string) *ModelRouterMiguUploadSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) SetSuccess(v bool) *ModelRouterMiguUploadSourceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterMiguUploadSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
