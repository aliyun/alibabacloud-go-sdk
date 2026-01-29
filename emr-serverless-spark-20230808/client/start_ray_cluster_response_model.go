// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *StartRayClusterResponseBody) *StartRayClusterResponse
	GetBody() *StartRayClusterResponseBody
}

type StartRayClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRayClusterResponse) GoString() string {
	return s.String()
}

func (s *StartRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRayClusterResponse) GetBody() *StartRayClusterResponseBody {
	return s.Body
}

func (s *StartRayClusterResponse) SetHeaders(v map[string]*string) *StartRayClusterResponse {
	s.Headers = v
	return s
}

func (s *StartRayClusterResponse) SetStatusCode(v int32) *StartRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRayClusterResponse) SetBody(v *StartRayClusterResponseBody) *StartRayClusterResponse {
	s.Body = v
	return s
}

func (s *StartRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
