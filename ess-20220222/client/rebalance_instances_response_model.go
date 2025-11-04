// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebalanceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebalanceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RebalanceInstancesResponseBody) *RebalanceInstancesResponse
	GetBody() *RebalanceInstancesResponseBody
}

type RebalanceInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebalanceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebalanceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebalanceInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebalanceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebalanceInstancesResponse) GetBody() *RebalanceInstancesResponseBody {
	return s.Body
}

func (s *RebalanceInstancesResponse) SetHeaders(v map[string]*string) *RebalanceInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebalanceInstancesResponse) SetStatusCode(v int32) *RebalanceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceInstancesResponse) SetBody(v *RebalanceInstancesResponseBody) *RebalanceInstancesResponse {
	s.Body = v
	return s
}

func (s *RebalanceInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
