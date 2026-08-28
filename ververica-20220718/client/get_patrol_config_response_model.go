// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPatrolConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPatrolConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPatrolConfigResponseBody) *GetPatrolConfigResponse
	GetBody() *GetPatrolConfigResponseBody
}

type GetPatrolConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPatrolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPatrolConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPatrolConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPatrolConfigResponse) GetBody() *GetPatrolConfigResponseBody {
	return s.Body
}

func (s *GetPatrolConfigResponse) SetHeaders(v map[string]*string) *GetPatrolConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPatrolConfigResponse) SetStatusCode(v int32) *GetPatrolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPatrolConfigResponse) SetBody(v *GetPatrolConfigResponseBody) *GetPatrolConfigResponse {
	s.Body = v
	return s
}

func (s *GetPatrolConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
