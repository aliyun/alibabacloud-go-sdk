// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuditTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAuditTaskResponseBody
	GetCode() *string
	SetData(v bool) *CancelAuditTaskResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CancelAuditTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelAuditTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAuditTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelAuditTaskResponseBody
	GetSuccess() *bool
}

type CancelAuditTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAuditTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAuditTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAuditTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAuditTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelAuditTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelAuditTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAuditTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAuditTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelAuditTaskResponseBody) SetCode(v string) *CancelAuditTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAuditTaskResponseBody) SetData(v bool) *CancelAuditTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelAuditTaskResponseBody) SetHttpStatusCode(v int32) *CancelAuditTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAuditTaskResponseBody) SetMessage(v string) *CancelAuditTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAuditTaskResponseBody) SetRequestId(v string) *CancelAuditTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAuditTaskResponseBody) SetSuccess(v bool) *CancelAuditTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAuditTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
