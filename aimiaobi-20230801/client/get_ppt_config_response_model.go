// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPptConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPptConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPptConfigResponseBody) *GetPptConfigResponse
	GetBody() *GetPptConfigResponseBody
}

type GetPptConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPptConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPptConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPptConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPptConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPptConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPptConfigResponse) GetBody() *GetPptConfigResponseBody {
	return s.Body
}

func (s *GetPptConfigResponse) SetHeaders(v map[string]*string) *GetPptConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPptConfigResponse) SetStatusCode(v int32) *GetPptConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPptConfigResponse) SetBody(v *GetPptConfigResponseBody) *GetPptConfigResponse {
	s.Body = v
	return s
}

func (s *GetPptConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
