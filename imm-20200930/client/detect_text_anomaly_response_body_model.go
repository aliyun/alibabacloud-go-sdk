// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTextAnomalyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetectTextAnomalyResponseBody
	GetRequestId() *string
	SetSuggestion(v string) *DetectTextAnomalyResponseBody
	GetSuggestion() *string
}

type DetectTextAnomalyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 91AC8C98-0F36-49D2-8290-742E24DF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the text contains anomalies. Valid values:
	//
	// 	- pass: the text does not contain anomalies.
	//
	// 	- block: the text contains anomalies.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DetectTextAnomalyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectTextAnomalyResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectTextAnomalyResponseBody) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DetectTextAnomalyResponseBody) SetRequestId(v string) *DetectTextAnomalyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectTextAnomalyResponseBody) SetSuggestion(v string) *DetectTextAnomalyResponseBody {
	s.Suggestion = &v
	return s
}

func (s *DetectTextAnomalyResponseBody) Validate() error {
	return dara.Validate(s)
}
