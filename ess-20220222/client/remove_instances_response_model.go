// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInstancesResponseBody) *RemoveInstancesResponse
	GetBody() *RemoveInstancesResponseBody
}

type RemoveInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInstancesResponse) GetBody() *RemoveInstancesResponseBody {
	return s.Body
}

func (s *RemoveInstancesResponse) SetHeaders(v map[string]*string) *RemoveInstancesResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstancesResponse) SetStatusCode(v int32) *RemoveInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstancesResponse) SetBody(v *RemoveInstancesResponseBody) *RemoveInstancesResponse {
	s.Body = v
	return s
}

func (s *RemoveInstancesResponse) Validate() error {
	return dara.Validate(s)
}
