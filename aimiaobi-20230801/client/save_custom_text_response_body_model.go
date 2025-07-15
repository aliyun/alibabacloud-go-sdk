// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveCustomTextResponseBody
	GetCode() *string
	SetData(v int64) *SaveCustomTextResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *SaveCustomTextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveCustomTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveCustomTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveCustomTextResponseBody
	GetSuccess() *bool
}

type SaveCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveCustomTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCustomTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveCustomTextResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SaveCustomTextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveCustomTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveCustomTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveCustomTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveCustomTextResponseBody) SetCode(v string) *SaveCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetData(v int64) *SaveCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetHttpStatusCode(v int32) *SaveCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetMessage(v string) *SaveCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetRequestId(v string) *SaveCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetSuccess(v bool) *SaveCustomTextResponseBody {
	s.Success = &v
	return s
}

func (s *SaveCustomTextResponseBody) Validate() error {
	return dara.Validate(s)
}
