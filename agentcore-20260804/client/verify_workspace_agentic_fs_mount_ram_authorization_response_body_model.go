// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetCode() *string
	SetData(v *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetData() *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData
	SetHttpStatusCode(v int32) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
	GetSuccess() *bool
}

type VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 550e8400-e29b-41d4-a716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetData() *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData {
	return s.Data
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetCode(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetData(v *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetHttpStatusCode(v int32) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetMessage(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetRequestId(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) SetSuccess(v bool) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData struct {
	// The authorization status. A value of AUTHORIZED does not indicate that the actual mount was successful.
	//
	// example:
	//
	// AUTHORIZED
	AuthorizationStatus *string `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) SetAuthorizationStatus(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData {
	s.AuthorizationStatus = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
