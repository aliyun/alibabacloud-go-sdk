// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshCRAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMeshCRAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMeshCRAggregationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMeshCRAggregationResponseBody) *UpdateMeshCRAggregationResponse
	GetBody() *UpdateMeshCRAggregationResponseBody
}

type UpdateMeshCRAggregationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeshCRAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMeshCRAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshCRAggregationResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMeshCRAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMeshCRAggregationResponse) GetBody() *UpdateMeshCRAggregationResponseBody {
	return s.Body
}

func (s *UpdateMeshCRAggregationResponse) SetHeaders(v map[string]*string) *UpdateMeshCRAggregationResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeshCRAggregationResponse) SetStatusCode(v int32) *UpdateMeshCRAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeshCRAggregationResponse) SetBody(v *UpdateMeshCRAggregationResponseBody) *UpdateMeshCRAggregationResponse {
	s.Body = v
	return s
}

func (s *UpdateMeshCRAggregationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
