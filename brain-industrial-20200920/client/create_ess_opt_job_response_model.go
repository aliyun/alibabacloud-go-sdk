// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEssOptJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEssOptJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEssOptJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateEssOptJobResponseBody) *CreateEssOptJobResponse
	GetBody() *CreateEssOptJobResponseBody
}

type CreateEssOptJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEssOptJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEssOptJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobResponse) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEssOptJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEssOptJobResponse) GetBody() *CreateEssOptJobResponseBody {
	return s.Body
}

func (s *CreateEssOptJobResponse) SetHeaders(v map[string]*string) *CreateEssOptJobResponse {
	s.Headers = v
	return s
}

func (s *CreateEssOptJobResponse) SetStatusCode(v int32) *CreateEssOptJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEssOptJobResponse) SetBody(v *CreateEssOptJobResponseBody) *CreateEssOptJobResponse {
	s.Body = v
	return s
}

func (s *CreateEssOptJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
