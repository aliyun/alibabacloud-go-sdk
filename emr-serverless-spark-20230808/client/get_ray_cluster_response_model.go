// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetRayClusterResponseBody) *GetRayClusterResponse
	GetBody() *GetRayClusterResponseBody
}

type GetRayClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRayClusterResponse) GoString() string {
	return s.String()
}

func (s *GetRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRayClusterResponse) GetBody() *GetRayClusterResponseBody {
	return s.Body
}

func (s *GetRayClusterResponse) SetHeaders(v map[string]*string) *GetRayClusterResponse {
	s.Headers = v
	return s
}

func (s *GetRayClusterResponse) SetStatusCode(v int32) *GetRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayClusterResponse) SetBody(v *GetRayClusterResponseBody) *GetRayClusterResponse {
	s.Body = v
	return s
}

func (s *GetRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
