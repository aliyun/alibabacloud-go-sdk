// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCollectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCollectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListCollectorsResponseBody) *ListCollectorsResponse
	GetBody() *ListCollectorsResponseBody
}

type ListCollectorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCollectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponse) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCollectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCollectorsResponse) GetBody() *ListCollectorsResponseBody {
	return s.Body
}

func (s *ListCollectorsResponse) SetHeaders(v map[string]*string) *ListCollectorsResponse {
	s.Headers = v
	return s
}

func (s *ListCollectorsResponse) SetStatusCode(v int32) *ListCollectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollectorsResponse) SetBody(v *ListCollectorsResponseBody) *ListCollectorsResponse {
	s.Body = v
	return s
}

func (s *ListCollectorsResponse) Validate() error {
	return dara.Validate(s)
}
