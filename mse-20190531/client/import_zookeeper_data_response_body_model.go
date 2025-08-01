// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportZookeeperDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ImportZookeeperDataResponseBody
	GetData() interface{}
	SetErrorCode(v string) *ImportZookeeperDataResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ImportZookeeperDataResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ImportZookeeperDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportZookeeperDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportZookeeperDataResponseBody
	GetSuccess() *bool
}

type ImportZookeeperDataResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// null
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 52BA6DA6-A702-4362-A32F-DFF79655****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportZookeeperDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportZookeeperDataResponseBody) GoString() string {
	return s.String()
}

func (s *ImportZookeeperDataResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ImportZookeeperDataResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ImportZookeeperDataResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ImportZookeeperDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportZookeeperDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportZookeeperDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportZookeeperDataResponseBody) SetData(v interface{}) *ImportZookeeperDataResponseBody {
	s.Data = v
	return s
}

func (s *ImportZookeeperDataResponseBody) SetErrorCode(v string) *ImportZookeeperDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportZookeeperDataResponseBody) SetHttpCode(v string) *ImportZookeeperDataResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ImportZookeeperDataResponseBody) SetMessage(v string) *ImportZookeeperDataResponseBody {
	s.Message = &v
	return s
}

func (s *ImportZookeeperDataResponseBody) SetRequestId(v string) *ImportZookeeperDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportZookeeperDataResponseBody) SetSuccess(v bool) *ImportZookeeperDataResponseBody {
	s.Success = &v
	return s
}

func (s *ImportZookeeperDataResponseBody) Validate() error {
	return dara.Validate(s)
}
