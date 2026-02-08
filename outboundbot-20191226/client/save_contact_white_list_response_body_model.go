// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedRows(v int32) *SaveContactWhiteListResponseBody
	GetAffectedRows() *int32
	SetCode(v string) *SaveContactWhiteListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveContactWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveContactWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveContactWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveContactWhiteListResponseBody
	GetSuccess() *bool
}

type SaveContactWhiteListResponseBody struct {
	// Number of affected rows
	//
	// example:
	//
	// 1
	AffectedRows *int32 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveContactWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContactWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContactWhiteListResponseBody) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *SaveContactWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveContactWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveContactWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveContactWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContactWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveContactWhiteListResponseBody) SetAffectedRows(v int32) *SaveContactWhiteListResponseBody {
	s.AffectedRows = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) SetCode(v string) *SaveContactWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) SetHttpStatusCode(v int32) *SaveContactWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) SetMessage(v string) *SaveContactWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) SetRequestId(v string) *SaveContactWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) SetSuccess(v bool) *SaveContactWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *SaveContactWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
