// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelFeatureResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelFeatureResponseBody) *CreateModelFeatureResponse
	GetBody() *CreateModelFeatureResponseBody
}

type CreateModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelFeatureResponse) GetBody() *CreateModelFeatureResponseBody {
	return s.Body
}

func (s *CreateModelFeatureResponse) SetHeaders(v map[string]*string) *CreateModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *CreateModelFeatureResponse) SetStatusCode(v int32) *CreateModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelFeatureResponse) SetBody(v *CreateModelFeatureResponseBody) *CreateModelFeatureResponse {
	s.Body = v
	return s
}

func (s *CreateModelFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
