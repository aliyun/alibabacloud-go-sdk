// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataSourceOrderConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveDataSourceOrderConfigResponseBody
	GetCode() *string
	SetData(v bool) *SaveDataSourceOrderConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SaveDataSourceOrderConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveDataSourceOrderConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveDataSourceOrderConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveDataSourceOrderConfigResponseBody
	GetSuccess() *bool
}

type SaveDataSourceOrderConfigResponseBody struct {
	// example:
	//
	// 200
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
	// successful
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

func (s SaveDataSourceOrderConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveDataSourceOrderConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveDataSourceOrderConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveDataSourceOrderConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveDataSourceOrderConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveDataSourceOrderConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveDataSourceOrderConfigResponseBody) SetCode(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetData(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *SaveDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetMessage(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetRequestId(v string) *SaveDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetSuccess(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
