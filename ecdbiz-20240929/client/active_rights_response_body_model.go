// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveRightsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActiveRightsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ActiveRightsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ActiveRightsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActiveRightsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActiveRightsResponseBody
	GetSuccess() *bool
}

type ActiveRightsResponseBody struct {
	// The business processing result code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The description of the processing result. This parameter is used to help troubleshoot failures.
	//
	// example:
	//
	// The request is processed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 01A00000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is processed. A successful activation request does not mean that the resource is created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveRightsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveRightsResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveRightsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActiveRightsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActiveRightsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActiveRightsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveRightsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveRightsResponseBody) SetCode(v string) *ActiveRightsResponseBody {
	s.Code = &v
	return s
}

func (s *ActiveRightsResponseBody) SetHttpStatusCode(v int32) *ActiveRightsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActiveRightsResponseBody) SetMessage(v string) *ActiveRightsResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveRightsResponseBody) SetRequestId(v string) *ActiveRightsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveRightsResponseBody) SetSuccess(v bool) *ActiveRightsResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveRightsResponseBody) Validate() error {
	return dara.Validate(s)
}
