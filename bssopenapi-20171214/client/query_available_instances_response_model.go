// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvailableInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvailableInstancesResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvailableInstancesResponseBody) *QueryAvailableInstancesResponse
	GetBody() *QueryAvailableInstancesResponseBody
}

type QueryAvailableInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailableInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvailableInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvailableInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvailableInstancesResponse) GetBody() *QueryAvailableInstancesResponseBody {
	return s.Body
}

func (s *QueryAvailableInstancesResponse) SetHeaders(v map[string]*string) *QueryAvailableInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableInstancesResponse) SetStatusCode(v int32) *QueryAvailableInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableInstancesResponse) SetBody(v *QueryAvailableInstancesResponseBody) *QueryAvailableInstancesResponse {
	s.Body = v
	return s
}

func (s *QueryAvailableInstancesResponse) Validate() error {
	return dara.Validate(s)
}
