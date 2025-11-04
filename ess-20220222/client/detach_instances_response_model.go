// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DetachInstancesResponseBody) *DetachInstancesResponse
	GetBody() *DetachInstancesResponseBody
}

type DetachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachInstancesResponse) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachInstancesResponse) GetBody() *DetachInstancesResponseBody {
	return s.Body
}

func (s *DetachInstancesResponse) SetHeaders(v map[string]*string) *DetachInstancesResponse {
	s.Headers = v
	return s
}

func (s *DetachInstancesResponse) SetStatusCode(v int32) *DetachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachInstancesResponse) SetBody(v *DetachInstancesResponseBody) *DetachInstancesResponse {
	s.Body = v
	return s
}

func (s *DetachInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
