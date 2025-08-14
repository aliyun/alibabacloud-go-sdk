// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRecommendTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateRecommendTaskResponseBody
	GetResultObject() *bool
}

type CreateRecommendTaskResponseBody struct {
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

func (s CreateRecommendTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecommendTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecommendTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateRecommendTaskResponseBody) SetRequestId(v string) *CreateRecommendTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecommendTaskResponseBody) SetResultObject(v bool) *CreateRecommendTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateRecommendTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
