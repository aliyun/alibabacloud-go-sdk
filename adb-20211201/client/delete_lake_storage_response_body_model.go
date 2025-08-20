// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLakeStorageResponseBody
	GetCode() *string
	SetData(v string) *DeleteLakeStorageResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteLakeStorageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteLakeStorageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLakeStorageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLakeStorageResponseBody
	GetSuccess() *bool
}

type DeleteLakeStorageResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the delete operation was successful.
	//
	// example:
	//
	// True
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the operation is asynchronously implemented, the job ID is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2******-B2C1-408E-AA73-DB8D59******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLakeStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLakeStorageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLakeStorageResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteLakeStorageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteLakeStorageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLakeStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLakeStorageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLakeStorageResponseBody) SetCode(v string) *DeleteLakeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) SetData(v string) *DeleteLakeStorageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) SetHttpStatusCode(v int32) *DeleteLakeStorageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) SetMessage(v string) *DeleteLakeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) SetRequestId(v string) *DeleteLakeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) SetSuccess(v bool) *DeleteLakeStorageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLakeStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
