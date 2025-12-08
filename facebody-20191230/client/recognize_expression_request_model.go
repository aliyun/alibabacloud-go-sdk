// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeExpressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeExpressionRequest
	GetImageURL() *string
}

type RecognizeExpressionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeExpression/RecognizeExpression10.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeExpressionRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeExpressionRequest) SetImageURL(v string) *RecognizeExpressionRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeExpressionRequest) Validate() error {
	return dara.Validate(s)
}
