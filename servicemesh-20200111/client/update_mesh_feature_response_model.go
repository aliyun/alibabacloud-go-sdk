// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMeshFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMeshFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMeshFeatureResponseBody) *UpdateMeshFeatureResponse
	GetBody() *UpdateMeshFeatureResponseBody
}

type UpdateMeshFeatureResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeshFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMeshFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMeshFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMeshFeatureResponse) GetBody() *UpdateMeshFeatureResponseBody {
	return s.Body
}

func (s *UpdateMeshFeatureResponse) SetHeaders(v map[string]*string) *UpdateMeshFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeshFeatureResponse) SetStatusCode(v int32) *UpdateMeshFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeshFeatureResponse) SetBody(v *UpdateMeshFeatureResponseBody) *UpdateMeshFeatureResponse {
	s.Body = v
	return s
}

func (s *UpdateMeshFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
