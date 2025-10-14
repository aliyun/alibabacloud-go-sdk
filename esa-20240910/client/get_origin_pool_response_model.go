// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginPoolResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginPoolResponseBody) *GetOriginPoolResponse
	GetBody() *GetOriginPoolResponseBody
}

type GetOriginPoolResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponse) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginPoolResponse) GetBody() *GetOriginPoolResponseBody {
	return s.Body
}

func (s *GetOriginPoolResponse) SetHeaders(v map[string]*string) *GetOriginPoolResponse {
	s.Headers = v
	return s
}

func (s *GetOriginPoolResponse) SetStatusCode(v int32) *GetOriginPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginPoolResponse) SetBody(v *GetOriginPoolResponseBody) *GetOriginPoolResponse {
	s.Body = v
	return s
}

func (s *GetOriginPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
