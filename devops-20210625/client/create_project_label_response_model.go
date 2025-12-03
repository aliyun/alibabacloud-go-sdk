// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectLabelResponse
	GetStatusCode() *int32
	SetBody(v *CreateProjectLabelResponseBody) *CreateProjectLabelResponse
	GetBody() *CreateProjectLabelResponseBody
}

type CreateProjectLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectLabelResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectLabelResponse) GetBody() *CreateProjectLabelResponseBody {
	return s.Body
}

func (s *CreateProjectLabelResponse) SetHeaders(v map[string]*string) *CreateProjectLabelResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectLabelResponse) SetStatusCode(v int32) *CreateProjectLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectLabelResponse) SetBody(v *CreateProjectLabelResponseBody) *CreateProjectLabelResponse {
	s.Body = v
	return s
}

func (s *CreateProjectLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
