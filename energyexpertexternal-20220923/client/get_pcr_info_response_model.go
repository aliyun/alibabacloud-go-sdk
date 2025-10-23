// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPcrInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPcrInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPcrInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPcrInfoResponseBody) *GetPcrInfoResponse
	GetBody() *GetPcrInfoResponseBody
}

type GetPcrInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPcrInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPcrInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPcrInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPcrInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPcrInfoResponse) GetBody() *GetPcrInfoResponseBody {
	return s.Body
}

func (s *GetPcrInfoResponse) SetHeaders(v map[string]*string) *GetPcrInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPcrInfoResponse) SetStatusCode(v int32) *GetPcrInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPcrInfoResponse) SetBody(v *GetPcrInfoResponseBody) *GetPcrInfoResponse {
	s.Body = v
	return s
}

func (s *GetPcrInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
