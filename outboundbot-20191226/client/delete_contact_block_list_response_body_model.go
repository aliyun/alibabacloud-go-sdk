// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactBlockListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedRows(v int32) *DeleteContactBlockListResponseBody
	GetAffectedRows() *int32
	SetCode(v string) *DeleteContactBlockListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteContactBlockListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteContactBlockListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactBlockListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactBlockListResponseBody
	GetSuccess() *bool
}

type DeleteContactBlockListResponseBody struct {
	// Number of affected rows [Deprecated]
	//
	// example:
	//
	// 1
	AffectedRows *int32 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// API status code
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

func (s DeleteContactBlockListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactBlockListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactBlockListResponseBody) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *DeleteContactBlockListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactBlockListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteContactBlockListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactBlockListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactBlockListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactBlockListResponseBody) SetAffectedRows(v int32) *DeleteContactBlockListResponseBody {
	s.AffectedRows = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) SetCode(v string) *DeleteContactBlockListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) SetHttpStatusCode(v int32) *DeleteContactBlockListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) SetMessage(v string) *DeleteContactBlockListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) SetRequestId(v string) *DeleteContactBlockListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) SetSuccess(v bool) *DeleteContactBlockListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactBlockListResponseBody) Validate() error {
	return dara.Validate(s)
}
