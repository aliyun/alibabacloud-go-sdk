// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncDepartmentResponseBody
	GetCode() *string
	SetData(v bool) *SyncDepartmentResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SyncDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SyncDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncDepartmentResponseBody
	GetSuccess() *bool
}

type SyncDepartmentResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncDepartmentResponseBody) GetData() *bool {
	return s.Data
}

func (s *SyncDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SyncDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncDepartmentResponseBody) SetCode(v string) *SyncDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *SyncDepartmentResponseBody) SetData(v bool) *SyncDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *SyncDepartmentResponseBody) SetHttpStatusCode(v int32) *SyncDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncDepartmentResponseBody) SetMessage(v string) *SyncDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDepartmentResponseBody) SetRequestId(v string) *SyncDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDepartmentResponseBody) SetSuccess(v bool) *SyncDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *SyncDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
