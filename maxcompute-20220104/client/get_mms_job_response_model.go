// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsJobResponseBody) *GetMmsJobResponse
	GetBody() *GetMmsJobResponseBody
}

type GetMmsJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsJobResponse) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsJobResponse) GetBody() *GetMmsJobResponseBody {
	return s.Body
}

func (s *GetMmsJobResponse) SetHeaders(v map[string]*string) *GetMmsJobResponse {
	s.Headers = v
	return s
}

func (s *GetMmsJobResponse) SetStatusCode(v int32) *GetMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsJobResponse) SetBody(v *GetMmsJobResponseBody) *GetMmsJobResponse {
	s.Body = v
	return s
}

func (s *GetMmsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
