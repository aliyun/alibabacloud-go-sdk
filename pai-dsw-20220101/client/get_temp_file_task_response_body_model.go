// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTempFileTaskResponseBody
	GetCode() *string
	SetGmtCreateTime(v string) *GetTempFileTaskResponseBody
	GetGmtCreateTime() *string
	SetGmtExpiredTime(v string) *GetTempFileTaskResponseBody
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *GetTempFileTaskResponseBody
	GetGmtModifiedTime() *string
	SetHttpStatusCode(v int32) *GetTempFileTaskResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetTempFileTaskResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetTempFileTaskResponseBody
	GetMessage() *string
	SetOwnerId(v string) *GetTempFileTaskResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *GetTempFileTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTempFileTaskResponseBody
	GetSuccess() *bool
	SetUserId(v string) *GetTempFileTaskResponseBody
	GetUserId() *string
	SetUuid(v string) *GetTempFileTaskResponseBody
	GetUuid() *string
}

type GetTempFileTaskResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Owner Id
	//
	// example:
	//
	// 1612285282502324
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1612285282502324
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// tempfile-05cexxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetTempFileTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTempFileTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTempFileTaskResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTempFileTaskResponseBody) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *GetTempFileTaskResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTempFileTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTempFileTaskResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTempFileTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTempFileTaskResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetTempFileTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTempFileTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTempFileTaskResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetTempFileTaskResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *GetTempFileTaskResponseBody) SetCode(v string) *GetTempFileTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetGmtCreateTime(v string) *GetTempFileTaskResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetGmtExpiredTime(v string) *GetTempFileTaskResponseBody {
	s.GmtExpiredTime = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetGmtModifiedTime(v string) *GetTempFileTaskResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetHttpStatusCode(v int32) *GetTempFileTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetInstanceId(v string) *GetTempFileTaskResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetMessage(v string) *GetTempFileTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetOwnerId(v string) *GetTempFileTaskResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetRequestId(v string) *GetTempFileTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetSuccess(v bool) *GetTempFileTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetUserId(v string) *GetTempFileTaskResponseBody {
	s.UserId = &v
	return s
}

func (s *GetTempFileTaskResponseBody) SetUuid(v string) *GetTempFileTaskResponseBody {
	s.Uuid = &v
	return s
}

func (s *GetTempFileTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
