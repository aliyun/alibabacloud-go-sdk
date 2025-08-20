// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApsManagedDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetApsManagedDatabasesResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetApsManagedDatabasesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApsManagedDatabasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApsManagedDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApsManagedDatabasesResponseBody
	GetSuccess() *bool
}

type GetApsManagedDatabasesResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
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

func (s GetApsManagedDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApsManagedDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetApsManagedDatabasesResponseBody) GetData() *string {
	return s.Data
}

func (s *GetApsManagedDatabasesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApsManagedDatabasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApsManagedDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApsManagedDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApsManagedDatabasesResponseBody) SetData(v string) *GetApsManagedDatabasesResponseBody {
	s.Data = &v
	return s
}

func (s *GetApsManagedDatabasesResponseBody) SetHttpStatusCode(v int32) *GetApsManagedDatabasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApsManagedDatabasesResponseBody) SetMessage(v string) *GetApsManagedDatabasesResponseBody {
	s.Message = &v
	return s
}

func (s *GetApsManagedDatabasesResponseBody) SetRequestId(v string) *GetApsManagedDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApsManagedDatabasesResponseBody) SetSuccess(v bool) *GetApsManagedDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *GetApsManagedDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}
