// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DuplicateModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DuplicateModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DuplicateModelResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DuplicateModelResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *DuplicateModelResponseBody
	GetSuccess() *bool
}

type DuplicateModelResponseBody struct {
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
	// 374
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DuplicateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DuplicateModelResponseBody) GoString() string {
	return s.String()
}

func (s *DuplicateModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DuplicateModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DuplicateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DuplicateModelResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DuplicateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DuplicateModelResponseBody) SetCode(v int64) *DuplicateModelResponseBody {
	s.Code = &v
	return s
}

func (s *DuplicateModelResponseBody) SetHttpStatusCode(v int64) *DuplicateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DuplicateModelResponseBody) SetRequestId(v string) *DuplicateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DuplicateModelResponseBody) SetResultObject(v string) *DuplicateModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DuplicateModelResponseBody) SetSuccess(v bool) *DuplicateModelResponseBody {
	s.Success = &v
	return s
}

func (s *DuplicateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
