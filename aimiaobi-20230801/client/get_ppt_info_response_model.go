// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPptInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPptInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPptInfoResponseBody) *GetPptInfoResponse
	GetBody() *GetPptInfoResponseBody
}

type GetPptInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPptInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPptInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPptInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPptInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPptInfoResponse) GetBody() *GetPptInfoResponseBody {
	return s.Body
}

func (s *GetPptInfoResponse) SetHeaders(v map[string]*string) *GetPptInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPptInfoResponse) SetStatusCode(v int32) *GetPptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPptInfoResponse) SetBody(v *GetPptInfoResponseBody) *GetPptInfoResponse {
	s.Body = v
	return s
}

func (s *GetPptInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
