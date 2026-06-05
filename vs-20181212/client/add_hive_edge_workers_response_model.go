// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHiveEdgeWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddHiveEdgeWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddHiveEdgeWorkersResponse
	GetStatusCode() *int32
	SetBody(v *AddHiveEdgeWorkersResponseBody) *AddHiveEdgeWorkersResponse
	GetBody() *AddHiveEdgeWorkersResponseBody
}

type AddHiveEdgeWorkersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHiveEdgeWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHiveEdgeWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersResponse) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddHiveEdgeWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddHiveEdgeWorkersResponse) GetBody() *AddHiveEdgeWorkersResponseBody {
	return s.Body
}

func (s *AddHiveEdgeWorkersResponse) SetHeaders(v map[string]*string) *AddHiveEdgeWorkersResponse {
	s.Headers = v
	return s
}

func (s *AddHiveEdgeWorkersResponse) SetStatusCode(v int32) *AddHiveEdgeWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHiveEdgeWorkersResponse) SetBody(v *AddHiveEdgeWorkersResponseBody) *AddHiveEdgeWorkersResponse {
	s.Body = v
	return s
}

func (s *AddHiveEdgeWorkersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
