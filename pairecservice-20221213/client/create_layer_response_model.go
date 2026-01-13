// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLayerResponse
	GetStatusCode() *int32
	SetBody(v *CreateLayerResponseBody) *CreateLayerResponse
	GetBody() *CreateLayerResponseBody
}

type CreateLayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerResponse) GoString() string {
	return s.String()
}

func (s *CreateLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLayerResponse) GetBody() *CreateLayerResponseBody {
	return s.Body
}

func (s *CreateLayerResponse) SetHeaders(v map[string]*string) *CreateLayerResponse {
	s.Headers = v
	return s
}

func (s *CreateLayerResponse) SetStatusCode(v int32) *CreateLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayerResponse) SetBody(v *CreateLayerResponseBody) *CreateLayerResponse {
	s.Body = v
	return s
}

func (s *CreateLayerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
