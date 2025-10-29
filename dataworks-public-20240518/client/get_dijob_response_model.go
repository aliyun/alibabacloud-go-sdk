// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDIJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDIJobResponseBody) *GetDIJobResponse
	GetBody() *GetDIJobResponseBody
}

type GetDIJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDIJobResponse) GetBody() *GetDIJobResponseBody {
	return s.Body
}

func (s *GetDIJobResponse) SetHeaders(v map[string]*string) *GetDIJobResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobResponse) SetStatusCode(v int32) *GetDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobResponse) SetBody(v *GetDIJobResponseBody) *GetDIJobResponse {
	s.Body = v
	return s
}

func (s *GetDIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
