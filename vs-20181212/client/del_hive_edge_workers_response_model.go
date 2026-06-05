// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelHiveEdgeWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelHiveEdgeWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelHiveEdgeWorkersResponse
	GetStatusCode() *int32
	SetBody(v *DelHiveEdgeWorkersResponseBody) *DelHiveEdgeWorkersResponse
	GetBody() *DelHiveEdgeWorkersResponseBody
}

type DelHiveEdgeWorkersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelHiveEdgeWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelHiveEdgeWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersResponse) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelHiveEdgeWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelHiveEdgeWorkersResponse) GetBody() *DelHiveEdgeWorkersResponseBody {
	return s.Body
}

func (s *DelHiveEdgeWorkersResponse) SetHeaders(v map[string]*string) *DelHiveEdgeWorkersResponse {
	s.Headers = v
	return s
}

func (s *DelHiveEdgeWorkersResponse) SetStatusCode(v int32) *DelHiveEdgeWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *DelHiveEdgeWorkersResponse) SetBody(v *DelHiveEdgeWorkersResponseBody) *DelHiveEdgeWorkersResponse {
	s.Body = v
	return s
}

func (s *DelHiveEdgeWorkersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
