// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAutoNormalReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAutoNormalReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAutoNormalReviewResponse
	GetStatusCode() *int32
	SetBody(v *OperationResult) *StopAutoNormalReviewResponse
	GetBody() *OperationResult
}

type StopAutoNormalReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAutoNormalReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAutoNormalReviewResponse) GoString() string {
	return s.String()
}

func (s *StopAutoNormalReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAutoNormalReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAutoNormalReviewResponse) GetBody() *OperationResult {
	return s.Body
}

func (s *StopAutoNormalReviewResponse) SetHeaders(v map[string]*string) *StopAutoNormalReviewResponse {
	s.Headers = v
	return s
}

func (s *StopAutoNormalReviewResponse) SetStatusCode(v int32) *StopAutoNormalReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAutoNormalReviewResponse) SetBody(v *OperationResult) *StopAutoNormalReviewResponse {
	s.Body = v
	return s
}

func (s *StopAutoNormalReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
