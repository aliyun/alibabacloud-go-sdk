// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllNumberDistrictInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAllNumberDistrictInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAllNumberDistrictInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAllNumberDistrictInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAllNumberDistrictInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAllNumberDistrictInfoResponseBody
	GetSuccess() *bool
}

type DeleteAllNumberDistrictInfoResponseBody struct {
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

func (s DeleteAllNumberDistrictInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllNumberDistrictInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllNumberDistrictInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAllNumberDistrictInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAllNumberDistrictInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAllNumberDistrictInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAllNumberDistrictInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAllNumberDistrictInfoResponseBody) SetCode(v string) *DeleteAllNumberDistrictInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponseBody) SetHttpStatusCode(v int32) *DeleteAllNumberDistrictInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponseBody) SetMessage(v string) *DeleteAllNumberDistrictInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponseBody) SetRequestId(v string) *DeleteAllNumberDistrictInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponseBody) SetSuccess(v bool) *DeleteAllNumberDistrictInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
