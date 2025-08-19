// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CompareFacesResponseBody
	GetCode() *string
	SetData(v *CompareFacesResponseBodyData) *CompareFacesResponseBody
	GetData() *CompareFacesResponseBodyData
	SetMessage(v string) *CompareFacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CompareFacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CompareFacesResponseBody
	GetSuccess() *bool
}

type CompareFacesResponseBody struct {
	// HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Result of the face comparison.
	Data *CompareFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// Error.InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the current request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CompareFacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompareFacesResponseBody) GetData() *CompareFacesResponseBodyData {
	return s.Data
}

func (s *CompareFacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompareFacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareFacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CompareFacesResponseBody) SetCode(v string) *CompareFacesResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFacesResponseBody) SetData(v *CompareFacesResponseBodyData) *CompareFacesResponseBody {
	s.Data = v
	return s
}

func (s *CompareFacesResponseBody) SetMessage(v string) *CompareFacesResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFacesResponseBody) SetRequestId(v string) *CompareFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponseBody) SetSuccess(v bool) *CompareFacesResponseBody {
	s.Success = &v
	return s
}

func (s *CompareFacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CompareFacesResponseBodyData struct {
	// Confidence thresholds for face comparison. The returned content is a JSON Object, with the specific structure being `"key":"value"`.
	//
	// - `key` represents the false acceptance rate, which is the probability of misidentifying someone else as the specified person.
	//
	// - `value` is the corresponding threshold.
	//
	//
	// > Regarding the confidence thresholds (confidenceThresholds) in the example:
	//
	// - `"0.0001": "90.07"` indicates that the threshold is 90.07 when the false acceptance rate is 0.01%.
	//
	// - `"0.001": "80.01"` indicates that the threshold is 80.01 when the false acceptance rate is 0.1%.
	//
	// - `"0.01": "70.02"` indicates that the threshold is 70.02 when the false acceptance rate is 1%.
	//
	// Confidence thresholds are dynamically provided based on different images and algorithms, so do not persist these thresholds.
	//
	// example:
	//
	// {"0.0001":"90.07","0.001":"80.01","0.01":"70.02"}
	ConfidenceThresholds *string `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty"`
	// The degree of similarity between the faces in the two images. The value range is [0, 100], with higher values indicating greater similarity.
	//
	// example:
	//
	// 98.7913
	SimilarityScore *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s CompareFacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompareFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBodyData) GetConfidenceThresholds() *string {
	return s.ConfidenceThresholds
}

func (s *CompareFacesResponseBodyData) GetSimilarityScore() *float32 {
	return s.SimilarityScore
}

func (s *CompareFacesResponseBodyData) SetConfidenceThresholds(v string) *CompareFacesResponseBodyData {
	s.ConfidenceThresholds = &v
	return s
}

func (s *CompareFacesResponseBodyData) SetSimilarityScore(v float32) *CompareFacesResponseBodyData {
	s.SimilarityScore = &v
	return s
}

func (s *CompareFacesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
