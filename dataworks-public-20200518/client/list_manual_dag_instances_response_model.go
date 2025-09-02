// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualDagInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManualDagInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManualDagInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListManualDagInstancesResponseBody) *ListManualDagInstancesResponse
	GetBody() *ListManualDagInstancesResponseBody
}

type ListManualDagInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManualDagInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManualDagInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManualDagInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListManualDagInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManualDagInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManualDagInstancesResponse) GetBody() *ListManualDagInstancesResponseBody {
	return s.Body
}

func (s *ListManualDagInstancesResponse) SetHeaders(v map[string]*string) *ListManualDagInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListManualDagInstancesResponse) SetStatusCode(v int32) *ListManualDagInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManualDagInstancesResponse) SetBody(v *ListManualDagInstancesResponseBody) *ListManualDagInstancesResponse {
	s.Body = v
	return s
}

func (s *ListManualDagInstancesResponse) Validate() error {
	return dara.Validate(s)
}
