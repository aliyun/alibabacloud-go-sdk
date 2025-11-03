// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPublishedRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcPublishedRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcPublishedRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcPublishedRouteEntriesResponseBody) *ListVpcPublishedRouteEntriesResponse
	GetBody() *ListVpcPublishedRouteEntriesResponseBody
}

type ListVpcPublishedRouteEntriesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcPublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcPublishedRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcPublishedRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcPublishedRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcPublishedRouteEntriesResponse) GetBody() *ListVpcPublishedRouteEntriesResponseBody {
	return s.Body
}

func (s *ListVpcPublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *ListVpcPublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponse) SetStatusCode(v int32) *ListVpcPublishedRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponse) SetBody(v *ListVpcPublishedRouteEntriesResponseBody) *ListVpcPublishedRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
