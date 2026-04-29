// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFeatureViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFeatureViewResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFeatureViewResponseBody) *UpdateFeatureViewResponse
	GetBody() *UpdateFeatureViewResponseBody
}

type UpdateFeatureViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeatureViewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFeatureViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFeatureViewResponse) GetBody() *UpdateFeatureViewResponseBody {
	return s.Body
}

func (s *UpdateFeatureViewResponse) SetHeaders(v map[string]*string) *UpdateFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeatureViewResponse) SetStatusCode(v int32) *UpdateFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeatureViewResponse) SetBody(v *UpdateFeatureViewResponseBody) *UpdateFeatureViewResponse {
	s.Body = v
	return s
}

func (s *UpdateFeatureViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
