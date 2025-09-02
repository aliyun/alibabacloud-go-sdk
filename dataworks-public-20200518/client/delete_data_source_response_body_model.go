// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteDataSourceResponseBody
	GetData() *bool
	SetHttpStatusCode(v string) *DeleteDataSourceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataSourceResponseBody
	GetSuccess() *bool
}

type DeleteDataSourceResponseBody struct {
	// Indicates whether the data source was removed.
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
	// 0bc1411515937635****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDataSourceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataSourceResponseBody) SetData(v bool) *DeleteDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetHttpStatusCode(v string) *DeleteDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetSuccess(v bool) *DeleteDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
