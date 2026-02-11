// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseExpressionParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ParseExpressionParametersResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *ParseExpressionParametersResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ParseExpressionParametersResponseBody
	GetRequestId() *string
	SetResultObject(v *ParseExpressionParametersResponseBodyResultObject) *ParseExpressionParametersResponseBody
	GetResultObject() *ParseExpressionParametersResponseBodyResultObject
	SetSuccess(v bool) *ParseExpressionParametersResponseBody
	GetSuccess() *bool
}

type ParseExpressionParametersResponseBody struct {
	// Status code. A return of 200 indicates success.
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
	ResultObject *ParseExpressionParametersResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ParseExpressionParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ParseExpressionParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ParseExpressionParametersResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ParseExpressionParametersResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ParseExpressionParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ParseExpressionParametersResponseBody) GetResultObject() *ParseExpressionParametersResponseBodyResultObject {
	return s.ResultObject
}

func (s *ParseExpressionParametersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ParseExpressionParametersResponseBody) SetCode(v int64) *ParseExpressionParametersResponseBody {
	s.Code = &v
	return s
}

func (s *ParseExpressionParametersResponseBody) SetHttpStatusCode(v int64) *ParseExpressionParametersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ParseExpressionParametersResponseBody) SetRequestId(v string) *ParseExpressionParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ParseExpressionParametersResponseBody) SetResultObject(v *ParseExpressionParametersResponseBodyResultObject) *ParseExpressionParametersResponseBody {
	s.ResultObject = v
	return s
}

func (s *ParseExpressionParametersResponseBody) SetSuccess(v bool) *ParseExpressionParametersResponseBody {
	s.Success = &v
	return s
}

func (s *ParseExpressionParametersResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ParseExpressionParametersResponseBodyResultObject struct {
	// List of parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s ParseExpressionParametersResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ParseExpressionParametersResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ParseExpressionParametersResponseBodyResultObject) GetParams() []*string {
	return s.Params
}

func (s *ParseExpressionParametersResponseBodyResultObject) SetParams(v []*string) *ParseExpressionParametersResponseBodyResultObject {
	s.Params = v
	return s
}

func (s *ParseExpressionParametersResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
