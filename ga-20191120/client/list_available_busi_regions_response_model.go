// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableBusiRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableBusiRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableBusiRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableBusiRegionsResponseBody) *ListAvailableBusiRegionsResponse
	GetBody() *ListAvailableBusiRegionsResponseBody
}

type ListAvailableBusiRegionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableBusiRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableBusiRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableBusiRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableBusiRegionsResponse) GetBody() *ListAvailableBusiRegionsResponseBody {
	return s.Body
}

func (s *ListAvailableBusiRegionsResponse) SetHeaders(v map[string]*string) *ListAvailableBusiRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableBusiRegionsResponse) SetStatusCode(v int32) *ListAvailableBusiRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableBusiRegionsResponse) SetBody(v *ListAvailableBusiRegionsResponseBody) *ListAvailableBusiRegionsResponse {
	s.Body = v
	return s
}

func (s *ListAvailableBusiRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
