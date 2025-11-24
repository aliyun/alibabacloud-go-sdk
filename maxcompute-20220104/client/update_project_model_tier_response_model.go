// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectModelTierResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectModelTierResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectModelTierResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectModelTierResponseBody) *UpdateProjectModelTierResponse
	GetBody() *UpdateProjectModelTierResponseBody
}

type UpdateProjectModelTierResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectModelTierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectModelTierResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectModelTierResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectModelTierResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectModelTierResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectModelTierResponse) GetBody() *UpdateProjectModelTierResponseBody {
	return s.Body
}

func (s *UpdateProjectModelTierResponse) SetHeaders(v map[string]*string) *UpdateProjectModelTierResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectModelTierResponse) SetStatusCode(v int32) *UpdateProjectModelTierResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectModelTierResponse) SetBody(v *UpdateProjectModelTierResponseBody) *UpdateProjectModelTierResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectModelTierResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
