// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSemanticJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSemanticJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateSemanticJobResponseBody) *CreateSemanticJobResponse
	GetBody() *CreateSemanticJobResponseBody
}

type CreateSemanticJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSemanticJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSemanticJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSemanticJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSemanticJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSemanticJobResponse) GetBody() *CreateSemanticJobResponseBody {
	return s.Body
}

func (s *CreateSemanticJobResponse) SetHeaders(v map[string]*string) *CreateSemanticJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSemanticJobResponse) SetStatusCode(v int32) *CreateSemanticJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSemanticJobResponse) SetBody(v *CreateSemanticJobResponseBody) *CreateSemanticJobResponse {
	s.Body = v
	return s
}

func (s *CreateSemanticJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
