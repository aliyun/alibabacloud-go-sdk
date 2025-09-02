// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateDataSourceResponseBody
	GetData() *bool
	SetHttpStatusCode(v string) *UpdateDataSourceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataSourceResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceResponseBody struct {
	// Indicates whether the data source is updated.
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
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115159376359****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDataSourceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceResponseBody) SetData(v bool) *UpdateDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetHttpStatusCode(v string) *UpdateDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
