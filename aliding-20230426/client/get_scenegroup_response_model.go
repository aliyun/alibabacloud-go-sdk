// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenegroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScenegroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScenegroupResponse
	GetStatusCode() *int32
	SetBody(v *GetScenegroupResponseBody) *GetScenegroupResponse
	GetBody() *GetScenegroupResponseBody
}

type GetScenegroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScenegroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScenegroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupResponse) GoString() string {
	return s.String()
}

func (s *GetScenegroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScenegroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScenegroupResponse) GetBody() *GetScenegroupResponseBody {
	return s.Body
}

func (s *GetScenegroupResponse) SetHeaders(v map[string]*string) *GetScenegroupResponse {
	s.Headers = v
	return s
}

func (s *GetScenegroupResponse) SetStatusCode(v int32) *GetScenegroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenegroupResponse) SetBody(v *GetScenegroupResponseBody) *GetScenegroupResponse {
	s.Body = v
	return s
}

func (s *GetScenegroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
