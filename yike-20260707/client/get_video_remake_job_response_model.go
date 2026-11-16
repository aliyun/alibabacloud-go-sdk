// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRemakeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoRemakeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoRemakeJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoRemakeJobResponseBody) *GetVideoRemakeJobResponse
	GetBody() *GetVideoRemakeJobResponseBody
}

type GetVideoRemakeJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoRemakeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoRemakeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRemakeJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoRemakeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoRemakeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoRemakeJobResponse) GetBody() *GetVideoRemakeJobResponseBody {
	return s.Body
}

func (s *GetVideoRemakeJobResponse) SetHeaders(v map[string]*string) *GetVideoRemakeJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoRemakeJobResponse) SetStatusCode(v int32) *GetVideoRemakeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoRemakeJobResponse) SetBody(v *GetVideoRemakeJobResponseBody) *GetVideoRemakeJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoRemakeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
