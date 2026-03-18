// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateBackupResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateBackupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateBackupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateBackupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateBackupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBackupResponseBody
	GetSuccess() *bool
}

type UpdateBackupResponseBody struct {
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateBackupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateBackupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateBackupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBackupResponseBody) SetData(v bool) *UpdateBackupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBackupResponseBody) SetErrCode(v string) *UpdateBackupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateBackupResponseBody) SetErrMessage(v string) *UpdateBackupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateBackupResponseBody) SetHttpStatusCode(v int32) *UpdateBackupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBackupResponseBody) SetRequestId(v string) *UpdateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupResponseBody) SetSuccess(v bool) *UpdateBackupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
