// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageShieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageShieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageShieldResponse
	GetStatusCode() *int32
	SetBody(v *GetPageShieldResponseBody) *GetPageShieldResponse
	GetBody() *GetPageShieldResponseBody
}

type GetPageShieldResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageShieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageShieldResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageShieldResponse) GoString() string {
	return s.String()
}

func (s *GetPageShieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageShieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageShieldResponse) GetBody() *GetPageShieldResponseBody {
	return s.Body
}

func (s *GetPageShieldResponse) SetHeaders(v map[string]*string) *GetPageShieldResponse {
	s.Headers = v
	return s
}

func (s *GetPageShieldResponse) SetStatusCode(v int32) *GetPageShieldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageShieldResponse) SetBody(v *GetPageShieldResponseBody) *GetPageShieldResponse {
	s.Body = v
	return s
}

func (s *GetPageShieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
