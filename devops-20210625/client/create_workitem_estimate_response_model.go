// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemEstimateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkitemEstimateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkitemEstimateResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkitemEstimateResponseBody) *CreateWorkitemEstimateResponse
	GetBody() *CreateWorkitemEstimateResponseBody
}

type CreateWorkitemEstimateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkitemEstimateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkitemEstimateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemEstimateResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkitemEstimateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkitemEstimateResponse) GetBody() *CreateWorkitemEstimateResponseBody {
	return s.Body
}

func (s *CreateWorkitemEstimateResponse) SetHeaders(v map[string]*string) *CreateWorkitemEstimateResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemEstimateResponse) SetStatusCode(v int32) *CreateWorkitemEstimateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemEstimateResponse) SetBody(v *CreateWorkitemEstimateResponseBody) *CreateWorkitemEstimateResponse {
	s.Body = v
	return s
}

func (s *CreateWorkitemEstimateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
