// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAutoNormalReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenAutoNormalReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenAutoNormalReviewResponse
	GetStatusCode() *int32
	SetBody(v *OperationResult) *OpenAutoNormalReviewResponse
	GetBody() *OperationResult
}

type OpenAutoNormalReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAutoNormalReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenAutoNormalReviewResponse) GoString() string {
	return s.String()
}

func (s *OpenAutoNormalReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenAutoNormalReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenAutoNormalReviewResponse) GetBody() *OperationResult {
	return s.Body
}

func (s *OpenAutoNormalReviewResponse) SetHeaders(v map[string]*string) *OpenAutoNormalReviewResponse {
	s.Headers = v
	return s
}

func (s *OpenAutoNormalReviewResponse) SetStatusCode(v int32) *OpenAutoNormalReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAutoNormalReviewResponse) SetBody(v *OperationResult) *OpenAutoNormalReviewResponse {
	s.Body = v
	return s
}

func (s *OpenAutoNormalReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
