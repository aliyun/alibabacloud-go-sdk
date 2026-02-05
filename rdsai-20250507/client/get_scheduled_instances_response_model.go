// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledInstancesResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledInstancesResponseBody) *GetScheduledInstancesResponse
	GetBody() *GetScheduledInstancesResponseBody
}

type GetScheduledInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledInstancesResponse) GetBody() *GetScheduledInstancesResponseBody {
	return s.Body
}

func (s *GetScheduledInstancesResponse) SetHeaders(v map[string]*string) *GetScheduledInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledInstancesResponse) SetStatusCode(v int32) *GetScheduledInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledInstancesResponse) SetBody(v *GetScheduledInstancesResponseBody) *GetScheduledInstancesResponse {
	s.Body = v
	return s
}

func (s *GetScheduledInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
