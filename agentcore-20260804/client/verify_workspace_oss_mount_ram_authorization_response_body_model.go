// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceOssMountRamAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetCode() *string
	SetData(v *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetData() *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData
	SetHttpStatusCode(v int32) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyWorkspaceOssMountRamAuthorizationResponseBody
	GetSuccess() *bool
}

type VerifyWorkspaceOssMountRamAuthorizationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetData() *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData {
	return s.Data
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetCode(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetData(v *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetHttpStatusCode(v int32) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetMessage(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetRequestId(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) SetSuccess(v bool) *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyWorkspaceOssMountRamAuthorizationResponseBodyData struct {
	// The OSS storage authorization status.
	//
	// example:
	//
	// AUTHORIZED
	AuthorizationStatus *string `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) SetAuthorizationStatus(v string) *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData {
	s.AuthorizationStatus = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
