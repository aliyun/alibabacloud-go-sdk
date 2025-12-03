// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiZoneClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiZoneClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiZoneClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiZoneClusterResponseBody) *CreateMultiZoneClusterResponse
	GetBody() *CreateMultiZoneClusterResponseBody
}

type CreateMultiZoneClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiZoneClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiZoneClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiZoneClusterResponse) GetBody() *CreateMultiZoneClusterResponseBody {
	return s.Body
}

func (s *CreateMultiZoneClusterResponse) SetHeaders(v map[string]*string) *CreateMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiZoneClusterResponse) SetStatusCode(v int32) *CreateMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiZoneClusterResponse) SetBody(v *CreateMultiZoneClusterResponseBody) *CreateMultiZoneClusterResponse {
	s.Body = v
	return s
}

func (s *CreateMultiZoneClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
