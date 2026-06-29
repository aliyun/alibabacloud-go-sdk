// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncDepartmentUserResponseBody
	GetCode() *string
	SetData(v bool) *SyncDepartmentUserResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SyncDepartmentUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SyncDepartmentUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncDepartmentUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncDepartmentUserResponseBody
	GetSuccess() *bool
}

type SyncDepartmentUserResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned for the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncDepartmentUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncDepartmentUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *SyncDepartmentUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SyncDepartmentUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncDepartmentUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDepartmentUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncDepartmentUserResponseBody) SetCode(v string) *SyncDepartmentUserResponseBody {
	s.Code = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) SetData(v bool) *SyncDepartmentUserResponseBody {
	s.Data = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) SetHttpStatusCode(v int32) *SyncDepartmentUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) SetMessage(v string) *SyncDepartmentUserResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) SetRequestId(v string) *SyncDepartmentUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) SetSuccess(v bool) *SyncDepartmentUserResponseBody {
	s.Success = &v
	return s
}

func (s *SyncDepartmentUserResponseBody) Validate() error {
	return dara.Validate(s)
}
