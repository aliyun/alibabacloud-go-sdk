// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnalysisConditionFavoriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAnalysisConditionFavoriteResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteAnalysisConditionFavoriteResponseBody
	GetResultObject() *bool
}

type DeleteAnalysisConditionFavoriteResponseBody struct {
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

func (s DeleteAnalysisConditionFavoriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnalysisConditionFavoriteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAnalysisConditionFavoriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAnalysisConditionFavoriteResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteAnalysisConditionFavoriteResponseBody) SetRequestId(v string) *DeleteAnalysisConditionFavoriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteResponseBody) SetResultObject(v bool) *DeleteAnalysisConditionFavoriteResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteResponseBody) Validate() error {
	return dara.Validate(s)
}
