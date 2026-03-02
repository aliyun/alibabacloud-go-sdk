// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePbcReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokePbcReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokePbcReviewResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *RevokePbcReviewResponse
	GetBody() *CatalogCommonResult
}

type RevokePbcReviewResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePbcReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokePbcReviewResponse) GoString() string {
	return s.String()
}

func (s *RevokePbcReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokePbcReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokePbcReviewResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *RevokePbcReviewResponse) SetHeaders(v map[string]*string) *RevokePbcReviewResponse {
	s.Headers = v
	return s
}

func (s *RevokePbcReviewResponse) SetStatusCode(v int32) *RevokePbcReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePbcReviewResponse) SetBody(v *CatalogCommonResult) *RevokePbcReviewResponse {
	s.Body = v
	return s
}

func (s *RevokePbcReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
