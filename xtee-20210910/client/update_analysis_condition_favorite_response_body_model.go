// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnalysisConditionFavoriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAnalysisConditionFavoriteResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateAnalysisConditionFavoriteResponseBody
	GetResultObject() *bool
}

type UpdateAnalysisConditionFavoriteResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s UpdateAnalysisConditionFavoriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnalysisConditionFavoriteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisConditionFavoriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAnalysisConditionFavoriteResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateAnalysisConditionFavoriteResponseBody) SetRequestId(v string) *UpdateAnalysisConditionFavoriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteResponseBody) SetResultObject(v bool) *UpdateAnalysisConditionFavoriteResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteResponseBody) Validate() error {
	return dara.Validate(s)
}
