// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSampleResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DeleteSampleResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DeleteSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSampleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteSampleResponseBody
	GetResultObject() *bool
}

type DeleteSampleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DeleteSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSampleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSampleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteSampleResponseBody) SetCode(v string) *DeleteSampleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSampleResponseBody) SetHttpStatusCode(v string) *DeleteSampleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSampleResponseBody) SetMessage(v string) *DeleteSampleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSampleResponseBody) SetRequestId(v string) *DeleteSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSampleResponseBody) SetResultObject(v bool) *DeleteSampleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteSampleResponseBody) Validate() error {
	return dara.Validate(s)
}
