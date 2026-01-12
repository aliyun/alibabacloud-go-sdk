// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindSvcMapBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindSvcMapBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindSvcMapBindResponse
	GetStatusCode() *int32
	SetBody(v *FindSvcMapBindResponseBody) *FindSvcMapBindResponse
	GetBody() *FindSvcMapBindResponseBody
}

type FindSvcMapBindResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindSvcMapBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindSvcMapBindResponse) String() string {
	return dara.Prettify(s)
}

func (s FindSvcMapBindResponse) GoString() string {
	return s.String()
}

func (s *FindSvcMapBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindSvcMapBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindSvcMapBindResponse) GetBody() *FindSvcMapBindResponseBody {
	return s.Body
}

func (s *FindSvcMapBindResponse) SetHeaders(v map[string]*string) *FindSvcMapBindResponse {
	s.Headers = v
	return s
}

func (s *FindSvcMapBindResponse) SetStatusCode(v int32) *FindSvcMapBindResponse {
	s.StatusCode = &v
	return s
}

func (s *FindSvcMapBindResponse) SetBody(v *FindSvcMapBindResponseBody) *FindSvcMapBindResponse {
	s.Body = v
	return s
}

func (s *FindSvcMapBindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
