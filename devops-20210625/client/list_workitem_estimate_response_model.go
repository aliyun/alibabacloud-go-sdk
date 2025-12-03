// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemEstimateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkitemEstimateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkitemEstimateResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkitemEstimateResponseBody) *ListWorkitemEstimateResponse
	GetBody() *ListWorkitemEstimateResponseBody
}

type ListWorkitemEstimateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkitemEstimateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkitemEstimateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemEstimateResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemEstimateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkitemEstimateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkitemEstimateResponse) GetBody() *ListWorkitemEstimateResponseBody {
	return s.Body
}

func (s *ListWorkitemEstimateResponse) SetHeaders(v map[string]*string) *ListWorkitemEstimateResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemEstimateResponse) SetStatusCode(v int32) *ListWorkitemEstimateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemEstimateResponse) SetBody(v *ListWorkitemEstimateResponseBody) *ListWorkitemEstimateResponse {
	s.Body = v
	return s
}

func (s *ListWorkitemEstimateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
