// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalAcceleratorResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGlobalAcceleratorResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGlobalAcceleratorResourcesResponse
	GetStatusCode() *int32
	SetBody(v *GetGlobalAcceleratorResourcesResponseBody) *GetGlobalAcceleratorResourcesResponse
	GetBody() *GetGlobalAcceleratorResourcesResponseBody
}

type GetGlobalAcceleratorResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGlobalAcceleratorResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGlobalAcceleratorResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalAcceleratorResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalAcceleratorResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGlobalAcceleratorResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGlobalAcceleratorResourcesResponse) GetBody() *GetGlobalAcceleratorResourcesResponseBody {
	return s.Body
}

func (s *GetGlobalAcceleratorResourcesResponse) SetHeaders(v map[string]*string) *GetGlobalAcceleratorResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponse) SetStatusCode(v int32) *GetGlobalAcceleratorResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponse) SetBody(v *GetGlobalAcceleratorResourcesResponseBody) *GetGlobalAcceleratorResourcesResponse {
	s.Body = v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
