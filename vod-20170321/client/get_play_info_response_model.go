// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlayInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPlayInfoResponseBody) *GetPlayInfoResponse
	GetBody() *GetPlayInfoResponseBody
}

type GetPlayInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlayInfoResponse) GetBody() *GetPlayInfoResponseBody {
	return s.Body
}

func (s *GetPlayInfoResponse) SetHeaders(v map[string]*string) *GetPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPlayInfoResponse) SetStatusCode(v int32) *GetPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlayInfoResponse) SetBody(v *GetPlayInfoResponseBody) *GetPlayInfoResponse {
	s.Body = v
	return s
}

func (s *GetPlayInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
