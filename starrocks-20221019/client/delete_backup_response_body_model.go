// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteBackupResponseBody
	GetData() *bool
	SetErrCode(v string) *DeleteBackupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteBackupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteBackupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteBackupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupResponseBody
	GetSuccess() *bool
}

type DeleteBackupResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteBackupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteBackupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteBackupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupResponseBody) SetData(v bool) *DeleteBackupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBackupResponseBody) SetErrCode(v string) *DeleteBackupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteBackupResponseBody) SetErrMessage(v string) *DeleteBackupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteBackupResponseBody) SetHttpStatusCode(v int32) *DeleteBackupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBackupResponseBody) SetRequestId(v string) *DeleteBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupResponseBody) SetSuccess(v bool) *DeleteBackupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
