// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsSlsADBJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsSlsADBJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsSlsADBJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsSlsADBJobResponseBody) *CreateApsSlsADBJobResponse
	GetBody() *CreateApsSlsADBJobResponseBody
}

type CreateApsSlsADBJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsSlsADBJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsSlsADBJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobResponse) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsSlsADBJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsSlsADBJobResponse) GetBody() *CreateApsSlsADBJobResponseBody {
	return s.Body
}

func (s *CreateApsSlsADBJobResponse) SetHeaders(v map[string]*string) *CreateApsSlsADBJobResponse {
	s.Headers = v
	return s
}

func (s *CreateApsSlsADBJobResponse) SetStatusCode(v int32) *CreateApsSlsADBJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsSlsADBJobResponse) SetBody(v *CreateApsSlsADBJobResponseBody) *CreateApsSlsADBJobResponse {
	s.Body = v
	return s
}

func (s *CreateApsSlsADBJobResponse) Validate() error {
	return dara.Validate(s)
}
