// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiSparkApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKyuubiSparkApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKyuubiSparkApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListKyuubiSparkApplicationsResponseBody) *ListKyuubiSparkApplicationsResponse
	GetBody() *ListKyuubiSparkApplicationsResponseBody
}

type ListKyuubiSparkApplicationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiSparkApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKyuubiSparkApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKyuubiSparkApplicationsResponse) GetBody() *ListKyuubiSparkApplicationsResponseBody {
	return s.Body
}

func (s *ListKyuubiSparkApplicationsResponse) SetHeaders(v map[string]*string) *ListKyuubiSparkApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponse) SetStatusCode(v int32) *ListKyuubiSparkApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponse) SetBody(v *ListKyuubiSparkApplicationsResponseBody) *ListKyuubiSparkApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
