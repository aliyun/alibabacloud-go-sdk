// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachInstancesResponse
	GetStatusCode() *int32
	SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse
	GetBody() *AttachInstancesResponseBody
}

type AttachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachInstancesResponse) GetBody() *AttachInstancesResponseBody {
	return s.Body
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]*string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetStatusCode(v int32) *AttachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

func (s *AttachInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
