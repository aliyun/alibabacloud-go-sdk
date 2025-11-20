// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstancesResponse
	GetStatusCode() *int32
	SetBody(v *GetInstancesResponseBody) *GetInstancesResponse
	GetBody() *GetInstancesResponseBody
}

type GetInstancesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstancesResponse) GetBody() *GetInstancesResponseBody {
	return s.Body
}

func (s *GetInstancesResponse) SetHeaders(v map[string]*string) *GetInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesResponse) SetStatusCode(v int32) *GetInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesResponse) SetBody(v *GetInstancesResponseBody) *GetInstancesResponse {
	s.Body = v
	return s
}

func (s *GetInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
