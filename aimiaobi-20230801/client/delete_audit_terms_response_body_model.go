// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditTermsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAuditTermsResponseBody
	GetCode() *string
	SetData(v bool) *DeleteAuditTermsResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteAuditTermsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAuditTermsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAuditTermsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAuditTermsResponseBody
	GetSuccess() *bool
}

type DeleteAuditTermsResponseBody struct {
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

func (s DeleteAuditTermsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditTermsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuditTermsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAuditTermsResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAuditTermsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAuditTermsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAuditTermsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuditTermsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAuditTermsResponseBody) SetCode(v string) *DeleteAuditTermsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) SetData(v bool) *DeleteAuditTermsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) SetHttpStatusCode(v int32) *DeleteAuditTermsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) SetMessage(v string) *DeleteAuditTermsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) SetRequestId(v string) *DeleteAuditTermsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) SetSuccess(v bool) *DeleteAuditTermsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAuditTermsResponseBody) Validate() error {
	return dara.Validate(s)
}
