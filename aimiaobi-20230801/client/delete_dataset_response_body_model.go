// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDatasetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDatasetResponseBody
	GetSuccess() *bool
}

type DeleteDatasetResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
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

func (s DeleteDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDatasetResponseBody) SetCode(v string) *DeleteDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatasetResponseBody) SetHttpStatusCode(v int32) *DeleteDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDatasetResponseBody) SetMessage(v string) *DeleteDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDatasetResponseBody) SetRequestId(v string) *DeleteDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetResponseBody) SetSuccess(v bool) *DeleteDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
