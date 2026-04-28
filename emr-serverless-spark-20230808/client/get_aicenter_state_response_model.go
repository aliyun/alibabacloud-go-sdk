// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICenterStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICenterStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICenterStateResponse
	GetStatusCode() *int32
	SetBody(v *GetAICenterStateResponseBody) *GetAICenterStateResponse
	GetBody() *GetAICenterStateResponseBody
}

type GetAICenterStateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICenterStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICenterStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICenterStateResponse) GoString() string {
	return s.String()
}

func (s *GetAICenterStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICenterStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICenterStateResponse) GetBody() *GetAICenterStateResponseBody {
	return s.Body
}

func (s *GetAICenterStateResponse) SetHeaders(v map[string]*string) *GetAICenterStateResponse {
	s.Headers = v
	return s
}

func (s *GetAICenterStateResponse) SetStatusCode(v int32) *GetAICenterStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICenterStateResponse) SetBody(v *GetAICenterStateResponseBody) *GetAICenterStateResponse {
	s.Body = v
	return s
}

func (s *GetAICenterStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
