// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFactAuditUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFactAuditUrlResponseBody
	GetCode() *string
	SetData(v []*string) *GetFactAuditUrlResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *GetFactAuditUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetFactAuditUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFactAuditUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFactAuditUrlResponseBody
	GetSuccess() *bool
}

type GetFactAuditUrlResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s GetFactAuditUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFactAuditUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetFactAuditUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFactAuditUrlResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetFactAuditUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFactAuditUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFactAuditUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFactAuditUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFactAuditUrlResponseBody) SetCode(v string) *GetFactAuditUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetFactAuditUrlResponseBody) SetData(v []*string) *GetFactAuditUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetFactAuditUrlResponseBody) SetHttpStatusCode(v int32) *GetFactAuditUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFactAuditUrlResponseBody) SetMessage(v string) *GetFactAuditUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetFactAuditUrlResponseBody) SetRequestId(v string) *GetFactAuditUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFactAuditUrlResponseBody) SetSuccess(v bool) *GetFactAuditUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetFactAuditUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
