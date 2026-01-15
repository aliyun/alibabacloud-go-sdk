// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepublishModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *PrepublishModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *PrepublishModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *PrepublishModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *PrepublishModelResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *PrepublishModelResponseBody
	GetSuccess() *bool
}

type PrepublishModelResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PrepublishModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrepublishModelResponseBody) GoString() string {
	return s.String()
}

func (s *PrepublishModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PrepublishModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *PrepublishModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrepublishModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *PrepublishModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PrepublishModelResponseBody) SetCode(v int64) *PrepublishModelResponseBody {
	s.Code = &v
	return s
}

func (s *PrepublishModelResponseBody) SetHttpStatusCode(v int64) *PrepublishModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PrepublishModelResponseBody) SetRequestId(v string) *PrepublishModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepublishModelResponseBody) SetResultObject(v bool) *PrepublishModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *PrepublishModelResponseBody) SetSuccess(v bool) *PrepublishModelResponseBody {
	s.Success = &v
	return s
}

func (s *PrepublishModelResponseBody) Validate() error {
	return dara.Validate(s)
}
