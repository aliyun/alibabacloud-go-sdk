// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *AttachRCInstancesResponseBody) *AttachRCInstancesResponse
	GetBody() *AttachRCInstancesResponseBody
}

type AttachRCInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachRCInstancesResponse) GetBody() *AttachRCInstancesResponseBody {
	return s.Body
}

func (s *AttachRCInstancesResponse) SetHeaders(v map[string]*string) *AttachRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachRCInstancesResponse) SetStatusCode(v int32) *AttachRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachRCInstancesResponse) SetBody(v *AttachRCInstancesResponseBody) *AttachRCInstancesResponse {
	s.Body = v
	return s
}

func (s *AttachRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
