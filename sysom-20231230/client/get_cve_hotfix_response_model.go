// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCveHotfixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCveHotfixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCveHotfixResponse
	GetStatusCode() *int32
	SetBody(v *GetCveHotfixResponseBody) *GetCveHotfixResponse
	GetBody() *GetCveHotfixResponseBody
}

type GetCveHotfixResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCveHotfixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCveHotfixResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixResponse) GoString() string {
	return s.String()
}

func (s *GetCveHotfixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCveHotfixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCveHotfixResponse) GetBody() *GetCveHotfixResponseBody {
	return s.Body
}

func (s *GetCveHotfixResponse) SetHeaders(v map[string]*string) *GetCveHotfixResponse {
	s.Headers = v
	return s
}

func (s *GetCveHotfixResponse) SetStatusCode(v int32) *GetCveHotfixResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCveHotfixResponse) SetBody(v *GetCveHotfixResponseBody) *GetCveHotfixResponse {
	s.Body = v
	return s
}

func (s *GetCveHotfixResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
