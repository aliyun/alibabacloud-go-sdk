// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDataNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalDataNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalDataNetworkResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalDataNetworkResponseBody) *CreateGlobalDataNetworkResponse
	GetBody() *CreateGlobalDataNetworkResponseBody
}

type CreateGlobalDataNetworkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalDataNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalDataNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDataNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalDataNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalDataNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalDataNetworkResponse) GetBody() *CreateGlobalDataNetworkResponseBody {
	return s.Body
}

func (s *CreateGlobalDataNetworkResponse) SetHeaders(v map[string]*string) *CreateGlobalDataNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalDataNetworkResponse) SetStatusCode(v int32) *CreateGlobalDataNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalDataNetworkResponse) SetBody(v *CreateGlobalDataNetworkResponseBody) *CreateGlobalDataNetworkResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalDataNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
