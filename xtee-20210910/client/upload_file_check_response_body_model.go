// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadFileCheckResponseBody
	GetRequestId() *string
	SetResultObject(v *UploadFileCheckResponseBodyResultObject) *UploadFileCheckResponseBody
	GetResultObject() *UploadFileCheckResponseBodyResultObject
}

type UploadFileCheckResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *UploadFileCheckResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s UploadFileCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadFileCheckResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadFileCheckResponseBody) GetResultObject() *UploadFileCheckResponseBodyResultObject {
	return s.ResultObject
}

func (s *UploadFileCheckResponseBody) SetRequestId(v string) *UploadFileCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFileCheckResponseBody) SetResultObject(v *UploadFileCheckResponseBodyResultObject) *UploadFileCheckResponseBody {
	s.ResultObject = v
	return s
}

func (s *UploadFileCheckResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadFileCheckResponseBodyResultObject struct {
	// Number of effective rows
	//
	// example:
	//
	// 100
	EffectiveNumber *int32 `json:"effectiveNumber,omitempty" xml:"effectiveNumber,omitempty"`
	// Valid sample data
	ResultList []*string `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
	// Total number of rows
	//
	// example:
	//
	// 100
	TotalNumber *int32 `json:"totalNumber,omitempty" xml:"totalNumber,omitempty"`
}

func (s UploadFileCheckResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s UploadFileCheckResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *UploadFileCheckResponseBodyResultObject) GetEffectiveNumber() *int32 {
	return s.EffectiveNumber
}

func (s *UploadFileCheckResponseBodyResultObject) GetResultList() []*string {
	return s.ResultList
}

func (s *UploadFileCheckResponseBodyResultObject) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *UploadFileCheckResponseBodyResultObject) SetEffectiveNumber(v int32) *UploadFileCheckResponseBodyResultObject {
	s.EffectiveNumber = &v
	return s
}

func (s *UploadFileCheckResponseBodyResultObject) SetResultList(v []*string) *UploadFileCheckResponseBodyResultObject {
	s.ResultList = v
	return s
}

func (s *UploadFileCheckResponseBodyResultObject) SetTotalNumber(v int32) *UploadFileCheckResponseBodyResultObject {
	s.TotalNumber = &v
	return s
}

func (s *UploadFileCheckResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
