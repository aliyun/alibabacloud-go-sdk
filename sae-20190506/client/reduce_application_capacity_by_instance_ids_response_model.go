// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceApplicationCapacityByInstanceIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReduceApplicationCapacityByInstanceIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReduceApplicationCapacityByInstanceIdsResponse
	GetStatusCode() *int32
	SetBody(v *ReduceApplicationCapacityByInstanceIdsResponseBody) *ReduceApplicationCapacityByInstanceIdsResponse
	GetBody() *ReduceApplicationCapacityByInstanceIdsResponseBody
}

type ReduceApplicationCapacityByInstanceIdsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReduceApplicationCapacityByInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) GetBody() *ReduceApplicationCapacityByInstanceIdsResponseBody {
	return s.Body
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetHeaders(v map[string]*string) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetStatusCode(v int32) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetBody(v *ReduceApplicationCapacityByInstanceIdsResponseBody) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.Body = v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) Validate() error {
	return dara.Validate(s)
}
