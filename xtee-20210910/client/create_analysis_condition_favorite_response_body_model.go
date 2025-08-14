// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisConditionFavoriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAnalysisConditionFavoriteResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateAnalysisConditionFavoriteResponseBody
	GetResultObject() *bool
}

type CreateAnalysisConditionFavoriteResponseBody struct {
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

func (s CreateAnalysisConditionFavoriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisConditionFavoriteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnalysisConditionFavoriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnalysisConditionFavoriteResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateAnalysisConditionFavoriteResponseBody) SetRequestId(v string) *CreateAnalysisConditionFavoriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteResponseBody) SetResultObject(v bool) *CreateAnalysisConditionFavoriteResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteResponseBody) Validate() error {
	return dara.Validate(s)
}
