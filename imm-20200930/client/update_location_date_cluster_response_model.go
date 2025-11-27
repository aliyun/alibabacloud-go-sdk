// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationDateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLocationDateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLocationDateClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLocationDateClusterResponseBody) *UpdateLocationDateClusterResponse
	GetBody() *UpdateLocationDateClusterResponseBody
}

type UpdateLocationDateClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLocationDateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLocationDateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationDateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLocationDateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLocationDateClusterResponse) GetBody() *UpdateLocationDateClusterResponseBody {
	return s.Body
}

func (s *UpdateLocationDateClusterResponse) SetHeaders(v map[string]*string) *UpdateLocationDateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocationDateClusterResponse) SetStatusCode(v int32) *UpdateLocationDateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocationDateClusterResponse) SetBody(v *UpdateLocationDateClusterResponseBody) *UpdateLocationDateClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateLocationDateClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
