// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEiamInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEiamInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEiamInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListEiamInstancesResponseBody) *ListEiamInstancesResponse
	GetBody() *ListEiamInstancesResponseBody
}

type ListEiamInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEiamInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEiamInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEiamInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEiamInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEiamInstancesResponse) GetBody() *ListEiamInstancesResponseBody {
	return s.Body
}

func (s *ListEiamInstancesResponse) SetHeaders(v map[string]*string) *ListEiamInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEiamInstancesResponse) SetStatusCode(v int32) *ListEiamInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEiamInstancesResponse) SetBody(v *ListEiamInstancesResponseBody) *ListEiamInstancesResponse {
	s.Body = v
	return s
}

func (s *ListEiamInstancesResponse) Validate() error {
	return dara.Validate(s)
}
