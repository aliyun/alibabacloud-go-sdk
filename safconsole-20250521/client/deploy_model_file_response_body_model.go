// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeployModelFileResponseBody
	GetCode() *int64
	SetRequestId(v string) *DeployModelFileResponseBody
	GetRequestId() *string
	SetResultObject(v int64) *DeployModelFileResponseBody
	GetResultObject() *int64
	SetSuccess(v bool) *DeployModelFileResponseBody
	GetSuccess() *bool
}

type DeployModelFileResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *int64 `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
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

func (s DeployModelFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployModelFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeployModelFileResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeployModelFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployModelFileResponseBody) GetResultObject() *int64 {
	return s.ResultObject
}

func (s *DeployModelFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployModelFileResponseBody) SetCode(v int64) *DeployModelFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeployModelFileResponseBody) SetRequestId(v string) *DeployModelFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployModelFileResponseBody) SetResultObject(v int64) *DeployModelFileResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeployModelFileResponseBody) SetSuccess(v bool) *DeployModelFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeployModelFileResponseBody) Validate() error {
	return dara.Validate(s)
}
