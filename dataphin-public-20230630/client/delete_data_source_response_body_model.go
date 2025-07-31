// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataSourceResponseBody
	GetCode() *string
	SetData(v bool) *DeleteDataSourceResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteDataSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataSourceResponseBody
	GetSuccess() *bool
}

type DeleteDataSourceResponseBody struct {
	// example:
	//
	// OK
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
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteDataSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataSourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDataSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataSourceResponseBody) SetCode(v string) *DeleteDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetData(v bool) *DeleteDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetHttpStatusCode(v int32) *DeleteDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetMessage(v string) *DeleteDataSourceResponseBody {
	s.Message = &v
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
