// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVcoRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVcoRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVcoRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateVcoRouteEntryResponseBody) *CreateVcoRouteEntryResponse
	GetBody() *CreateVcoRouteEntryResponseBody
}

type CreateVcoRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVcoRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVcoRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVcoRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVcoRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVcoRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVcoRouteEntryResponse) GetBody() *CreateVcoRouteEntryResponseBody {
	return s.Body
}

func (s *CreateVcoRouteEntryResponse) SetHeaders(v map[string]*string) *CreateVcoRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateVcoRouteEntryResponse) SetStatusCode(v int32) *CreateVcoRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVcoRouteEntryResponse) SetBody(v *CreateVcoRouteEntryResponseBody) *CreateVcoRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateVcoRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
