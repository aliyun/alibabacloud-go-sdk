// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataServiceApiTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SaveDataServiceApiTestResultResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SaveDataServiceApiTestResultResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SaveDataServiceApiTestResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveDataServiceApiTestResultResponseBody
	GetSuccess() *bool
}

type SaveDataServiceApiTestResultResponseBody struct {
	// Indicates whether the test results are saved.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveDataServiceApiTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveDataServiceApiTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDataServiceApiTestResultResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveDataServiceApiTestResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveDataServiceApiTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveDataServiceApiTestResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveDataServiceApiTestResultResponseBody) SetData(v bool) *SaveDataServiceApiTestResultResponseBody {
	s.Data = &v
	return s
}

func (s *SaveDataServiceApiTestResultResponseBody) SetHttpStatusCode(v int32) *SaveDataServiceApiTestResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveDataServiceApiTestResultResponseBody) SetRequestId(v string) *SaveDataServiceApiTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDataServiceApiTestResultResponseBody) SetSuccess(v bool) *SaveDataServiceApiTestResultResponseBody {
	s.Success = &v
	return s
}

func (s *SaveDataServiceApiTestResultResponseBody) Validate() error {
	return dara.Validate(s)
}
