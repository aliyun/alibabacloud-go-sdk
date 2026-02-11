// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafRmmpOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeSafRmmpOrderResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeSafRmmpOrderResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeSafRmmpOrderResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSafRmmpOrderResponseBodyResultObject) *DescribeSafRmmpOrderResponseBody
	GetResultObject() *DescribeSafRmmpOrderResponseBodyResultObject
	SetSuccess(v bool) *DescribeSafRmmpOrderResponseBody
	GetSuccess() *bool
}

type DescribeSafRmmpOrderResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject *DescribeSafRmmpOrderResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSafRmmpOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafRmmpOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafRmmpOrderResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeSafRmmpOrderResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeSafRmmpOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafRmmpOrderResponseBody) GetResultObject() *DescribeSafRmmpOrderResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafRmmpOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSafRmmpOrderResponseBody) SetCode(v int64) *DescribeSafRmmpOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSafRmmpOrderResponseBody) SetHttpStatusCode(v int64) *DescribeSafRmmpOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSafRmmpOrderResponseBody) SetRequestId(v string) *DescribeSafRmmpOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafRmmpOrderResponseBody) SetResultObject(v *DescribeSafRmmpOrderResponseBodyResultObject) *DescribeSafRmmpOrderResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafRmmpOrderResponseBody) SetSuccess(v bool) *DescribeSafRmmpOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSafRmmpOrderResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSafRmmpOrderResponseBodyResultObject struct {
	// Indicates whether there is functional permission.
	//
	// example:
	//
	// true
	Accessible *bool `json:"Accessible,omitempty" xml:"Accessible,omitempty"`
}

func (s DescribeSafRmmpOrderResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafRmmpOrderResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafRmmpOrderResponseBodyResultObject) GetAccessible() *bool {
	return s.Accessible
}

func (s *DescribeSafRmmpOrderResponseBodyResultObject) SetAccessible(v bool) *DescribeSafRmmpOrderResponseBodyResultObject {
	s.Accessible = &v
	return s
}

func (s *DescribeSafRmmpOrderResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
