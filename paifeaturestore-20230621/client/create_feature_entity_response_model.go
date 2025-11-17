// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFeatureEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFeatureEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateFeatureEntityResponseBody) *CreateFeatureEntityResponse
	GetBody() *CreateFeatureEntityResponseBody
}

type CreateFeatureEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFeatureEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFeatureEntityResponse) GetBody() *CreateFeatureEntityResponseBody {
	return s.Body
}

func (s *CreateFeatureEntityResponse) SetHeaders(v map[string]*string) *CreateFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureEntityResponse) SetStatusCode(v int32) *CreateFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureEntityResponse) SetBody(v *CreateFeatureEntityResponseBody) *CreateFeatureEntityResponse {
	s.Body = v
	return s
}

func (s *CreateFeatureEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
