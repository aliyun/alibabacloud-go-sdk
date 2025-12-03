// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkitemFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkitemFileResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkitemFileResponseBody) *GetWorkitemFileResponse
	GetBody() *GetWorkitemFileResponseBody
}

type GetWorkitemFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkitemFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkitemFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemFileResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkitemFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkitemFileResponse) GetBody() *GetWorkitemFileResponseBody {
	return s.Body
}

func (s *GetWorkitemFileResponse) SetHeaders(v map[string]*string) *GetWorkitemFileResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemFileResponse) SetStatusCode(v int32) *GetWorkitemFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemFileResponse) SetBody(v *GetWorkitemFileResponseBody) *GetWorkitemFileResponse {
	s.Body = v
	return s
}

func (s *GetWorkitemFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
