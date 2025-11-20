// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveFormRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveFormRemarkResponse
	GetStatusCode() *int32
	SetBody(v *SaveFormRemarkResponseBody) *SaveFormRemarkResponse
	GetBody() *SaveFormRemarkResponseBody
}

type SaveFormRemarkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFormRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveFormRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkResponse) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveFormRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveFormRemarkResponse) GetBody() *SaveFormRemarkResponseBody {
	return s.Body
}

func (s *SaveFormRemarkResponse) SetHeaders(v map[string]*string) *SaveFormRemarkResponse {
	s.Headers = v
	return s
}

func (s *SaveFormRemarkResponse) SetStatusCode(v int32) *SaveFormRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFormRemarkResponse) SetBody(v *SaveFormRemarkResponseBody) *SaveFormRemarkResponse {
	s.Body = v
	return s
}

func (s *SaveFormRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
