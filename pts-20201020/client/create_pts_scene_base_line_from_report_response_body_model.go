// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneBaseLineFromReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePtsSceneBaseLineFromReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreatePtsSceneBaseLineFromReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePtsSceneBaseLineFromReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePtsSceneBaseLineFromReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePtsSceneBaseLineFromReportResponseBody
	GetSuccess() *bool
}

type CreatePtsSceneBaseLineFromReportResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F7D2CE0-AE4C-4143-954A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false:
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePtsSceneBaseLineFromReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetCode(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetHttpStatusCode(v int32) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetMessage(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetRequestId(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetSuccess(v bool) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) Validate() error {
	return dara.Validate(s)
}
