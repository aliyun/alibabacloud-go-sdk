// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFactAuditUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFactAuditUrlResponseBody
	GetCode() *string
	SetData(v string) *DeleteFactAuditUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteFactAuditUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFactAuditUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFactAuditUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFactAuditUrlResponseBody
	GetSuccess() *bool
}

type DeleteFactAuditUrlResponseBody struct {
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

func (s DeleteFactAuditUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFactAuditUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFactAuditUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFactAuditUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteFactAuditUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFactAuditUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFactAuditUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFactAuditUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFactAuditUrlResponseBody) SetCode(v string) *DeleteFactAuditUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) SetData(v string) *DeleteFactAuditUrlResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) SetHttpStatusCode(v int32) *DeleteFactAuditUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) SetMessage(v string) *DeleteFactAuditUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) SetRequestId(v string) *DeleteFactAuditUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) SetSuccess(v bool) *DeleteFactAuditUrlResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFactAuditUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
