// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAddonResponse
	GetStatusCode() *int32
	SetBody(v *GetAddonResponseBody) *GetAddonResponse
	GetBody() *GetAddonResponseBody
}

type GetAddonResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponse) GoString() string {
	return s.String()
}

func (s *GetAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAddonResponse) GetBody() *GetAddonResponseBody {
	return s.Body
}

func (s *GetAddonResponse) SetHeaders(v map[string]*string) *GetAddonResponse {
	s.Headers = v
	return s
}

func (s *GetAddonResponse) SetStatusCode(v int32) *GetAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddonResponse) SetBody(v *GetAddonResponseBody) *GetAddonResponse {
	s.Body = v
	return s
}

func (s *GetAddonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
