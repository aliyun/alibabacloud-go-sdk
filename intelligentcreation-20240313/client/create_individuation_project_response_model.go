// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIndividuationProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIndividuationProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateIndividuationProjectResponseBody) *CreateIndividuationProjectResponse
	GetBody() *CreateIndividuationProjectResponseBody
}

type CreateIndividuationProjectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndividuationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndividuationProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIndividuationProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIndividuationProjectResponse) GetBody() *CreateIndividuationProjectResponseBody {
	return s.Body
}

func (s *CreateIndividuationProjectResponse) SetHeaders(v map[string]*string) *CreateIndividuationProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateIndividuationProjectResponse) SetStatusCode(v int32) *CreateIndividuationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndividuationProjectResponse) SetBody(v *CreateIndividuationProjectResponseBody) *CreateIndividuationProjectResponse {
	s.Body = v
	return s
}

func (s *CreateIndividuationProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
