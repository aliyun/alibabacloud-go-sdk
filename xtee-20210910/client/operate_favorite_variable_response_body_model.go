// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFavoriteVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateFavoriteVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *OperateFavoriteVariableResponseBody
	GetResultObject() *bool
}

type OperateFavoriteVariableResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s OperateFavoriteVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateFavoriteVariableResponseBody) GoString() string {
	return s.String()
}

func (s *OperateFavoriteVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateFavoriteVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *OperateFavoriteVariableResponseBody) SetRequestId(v string) *OperateFavoriteVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateFavoriteVariableResponseBody) SetResultObject(v bool) *OperateFavoriteVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *OperateFavoriteVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
