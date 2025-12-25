// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConnDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConnDataResponse
	GetStatusCode() *int32
	SetBody(v *GetConnDataResponseBody) *GetConnDataResponse
	GetBody() *GetConnDataResponseBody
}

type GetConnDataResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConnDataResponse) GoString() string {
	return s.String()
}

func (s *GetConnDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConnDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConnDataResponse) GetBody() *GetConnDataResponseBody {
	return s.Body
}

func (s *GetConnDataResponse) SetHeaders(v map[string]*string) *GetConnDataResponse {
	s.Headers = v
	return s
}

func (s *GetConnDataResponse) SetStatusCode(v int32) *GetConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnDataResponse) SetBody(v *GetConnDataResponseBody) *GetConnDataResponse {
	s.Body = v
	return s
}

func (s *GetConnDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
