// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedRows(v int32) *DeleteContactWhiteListResponseBody
	GetAffectedRows() *int32
	SetCode(v string) *DeleteContactWhiteListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteContactWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteContactWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactWhiteListResponseBody
	GetSuccess() *bool
}

type DeleteContactWhiteListResponseBody struct {
	// Number of affected rows [Deprecated]
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

func (s DeleteContactWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactWhiteListResponseBody) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *DeleteContactWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteContactWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactWhiteListResponseBody) SetAffectedRows(v int32) *DeleteContactWhiteListResponseBody {
	s.AffectedRows = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) SetCode(v string) *DeleteContactWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) SetHttpStatusCode(v int32) *DeleteContactWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) SetMessage(v string) *DeleteContactWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) SetRequestId(v string) *DeleteContactWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) SetSuccess(v bool) *DeleteContactWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
