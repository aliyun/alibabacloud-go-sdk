// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRecommendationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UploadRecommendationDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadRecommendationDataResponseBody
	GetRequestId() *string
}

type UploadRecommendationDataResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRecommendationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadRecommendationDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadRecommendationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadRecommendationDataResponseBody) SetMessage(v string) *UploadRecommendationDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRecommendationDataResponseBody) SetRequestId(v string) *UploadRecommendationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRecommendationDataResponseBody) Validate() error {
	return dara.Validate(s)
}
