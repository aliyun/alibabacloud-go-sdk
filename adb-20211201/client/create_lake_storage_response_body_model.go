// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLakeStorageResponseBody
	GetCode() *string
	SetData(v string) *CreateLakeStorageResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateLakeStorageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateLakeStorageResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLakeStorageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLakeStorageResponseBody
	GetSuccess() *bool
}

type CreateLakeStorageResponseBody struct {
	// The API status or POP error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The resource usage details of the cluster.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If the creation operation is implemented asynchronously, a specific JobId is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLakeStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLakeStorageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLakeStorageResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateLakeStorageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateLakeStorageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLakeStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLakeStorageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLakeStorageResponseBody) SetCode(v string) *CreateLakeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLakeStorageResponseBody) SetData(v string) *CreateLakeStorageResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLakeStorageResponseBody) SetHttpStatusCode(v int32) *CreateLakeStorageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateLakeStorageResponseBody) SetMessage(v string) *CreateLakeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLakeStorageResponseBody) SetRequestId(v string) *CreateLakeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLakeStorageResponseBody) SetSuccess(v bool) *CreateLakeStorageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLakeStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
