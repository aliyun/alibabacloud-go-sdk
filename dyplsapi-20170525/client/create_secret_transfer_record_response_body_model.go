// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretTransferRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSecretTransferRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSecretTransferRecordResponseBody
	GetCode() *string
	SetData(v int64) *CreateSecretTransferRecordResponseBody
	GetData() *int64
	SetMessage(v string) *CreateSecretTransferRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecretTransferRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecretTransferRecordResponseBody
	GetSuccess() *bool
}

type CreateSecretTransferRecordResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 38
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 123456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecretTransferRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretTransferRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretTransferRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSecretTransferRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecretTransferRecordResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSecretTransferRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecretTransferRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretTransferRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecretTransferRecordResponseBody) SetAccessDeniedDetail(v string) *CreateSecretTransferRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) SetCode(v string) *CreateSecretTransferRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) SetData(v int64) *CreateSecretTransferRecordResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) SetMessage(v string) *CreateSecretTransferRecordResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) SetRequestId(v string) *CreateSecretTransferRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) SetSuccess(v bool) *CreateSecretTransferRecordResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecretTransferRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
