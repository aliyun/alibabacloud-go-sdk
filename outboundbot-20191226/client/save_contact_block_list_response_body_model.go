// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactBlockListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedRows(v int32) *SaveContactBlockListResponseBody
	GetAffectedRows() *int32
	SetCode(v string) *SaveContactBlockListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveContactBlockListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveContactBlockListResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveContactBlockListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveContactBlockListResponseBody
	GetSuccess() *bool
}

type SaveContactBlockListResponseBody struct {
	// example:
	//
	// 5
	AffectedRows *int32 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveContactBlockListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContactBlockListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContactBlockListResponseBody) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *SaveContactBlockListResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveContactBlockListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveContactBlockListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveContactBlockListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContactBlockListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveContactBlockListResponseBody) SetAffectedRows(v int32) *SaveContactBlockListResponseBody {
	s.AffectedRows = &v
	return s
}

func (s *SaveContactBlockListResponseBody) SetCode(v string) *SaveContactBlockListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveContactBlockListResponseBody) SetHttpStatusCode(v int32) *SaveContactBlockListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveContactBlockListResponseBody) SetMessage(v string) *SaveContactBlockListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveContactBlockListResponseBody) SetRequestId(v string) *SaveContactBlockListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContactBlockListResponseBody) SetSuccess(v bool) *SaveContactBlockListResponseBody {
	s.Success = &v
	return s
}

func (s *SaveContactBlockListResponseBody) Validate() error {
	return dara.Validate(s)
}
