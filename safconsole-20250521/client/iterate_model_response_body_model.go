// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIterateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *IterateModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *IterateModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *IterateModelResponseBody
	GetRequestId() *string
	SetResultObject(v string) *IterateModelResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *IterateModelResponseBody
	GetSuccess() *bool
}

type IterateModelResponseBody struct {
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
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IterateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IterateModelResponseBody) GoString() string {
	return s.String()
}

func (s *IterateModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *IterateModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *IterateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IterateModelResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *IterateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IterateModelResponseBody) SetCode(v int64) *IterateModelResponseBody {
	s.Code = &v
	return s
}

func (s *IterateModelResponseBody) SetHttpStatusCode(v int64) *IterateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IterateModelResponseBody) SetRequestId(v string) *IterateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *IterateModelResponseBody) SetResultObject(v string) *IterateModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *IterateModelResponseBody) SetSuccess(v bool) *IterateModelResponseBody {
	s.Success = &v
	return s
}

func (s *IterateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
