// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateDataSourceResponseBody
	GetData() *int64
	SetHttpStatusCode(v string) *CreateDataSourceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataSourceResponseBody
	GetSuccess() *bool
}

type CreateDataSourceResponseBody struct {
	// The data source ID.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc141151593763****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDataSourceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataSourceResponseBody) SetData(v int64) *CreateDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetHttpStatusCode(v string) *CreateDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetSuccess(v bool) *CreateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
