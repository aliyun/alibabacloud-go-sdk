// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRiskCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRiskCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRiskCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRiskCheckTaskResponseBody) *CreateRiskCheckTaskResponse
	GetBody() *CreateRiskCheckTaskResponseBody
}

type CreateRiskCheckTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRiskCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRiskCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRiskCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRiskCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRiskCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRiskCheckTaskResponse) GetBody() *CreateRiskCheckTaskResponseBody {
	return s.Body
}

func (s *CreateRiskCheckTaskResponse) SetHeaders(v map[string]*string) *CreateRiskCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRiskCheckTaskResponse) SetStatusCode(v int32) *CreateRiskCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRiskCheckTaskResponse) SetBody(v *CreateRiskCheckTaskResponseBody) *CreateRiskCheckTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRiskCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
