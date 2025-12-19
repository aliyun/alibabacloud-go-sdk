// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPotentialFailZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPotentialFailZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPotentialFailZonesResponse
	GetStatusCode() *int32
	SetBody(v *GetPotentialFailZonesResponseBody) *GetPotentialFailZonesResponse
	GetBody() *GetPotentialFailZonesResponseBody
}

type GetPotentialFailZonesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPotentialFailZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPotentialFailZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPotentialFailZonesResponse) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPotentialFailZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPotentialFailZonesResponse) GetBody() *GetPotentialFailZonesResponseBody {
	return s.Body
}

func (s *GetPotentialFailZonesResponse) SetHeaders(v map[string]*string) *GetPotentialFailZonesResponse {
	s.Headers = v
	return s
}

func (s *GetPotentialFailZonesResponse) SetStatusCode(v int32) *GetPotentialFailZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPotentialFailZonesResponse) SetBody(v *GetPotentialFailZonesResponseBody) *GetPotentialFailZonesResponse {
	s.Body = v
	return s
}

func (s *GetPotentialFailZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
