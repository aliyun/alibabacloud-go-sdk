// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFactAuditUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitFactAuditUrlResponseBody
	GetCode() *string
	SetData(v string) *SubmitFactAuditUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *SubmitFactAuditUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitFactAuditUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitFactAuditUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitFactAuditUrlResponseBody
	GetSuccess() *bool
}

type SubmitFactAuditUrlResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESSED
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitFactAuditUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitFactAuditUrlResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFactAuditUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitFactAuditUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitFactAuditUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitFactAuditUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitFactAuditUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitFactAuditUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitFactAuditUrlResponseBody) SetCode(v string) *SubmitFactAuditUrlResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) SetData(v string) *SubmitFactAuditUrlResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) SetHttpStatusCode(v int32) *SubmitFactAuditUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) SetMessage(v string) *SubmitFactAuditUrlResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) SetRequestId(v string) *SubmitFactAuditUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) SetSuccess(v bool) *SubmitFactAuditUrlResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitFactAuditUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
