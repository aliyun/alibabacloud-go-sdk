// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorFiltersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficMirrorFiltersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficMirrorFiltersResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficMirrorFiltersResponseBody) *ListTrafficMirrorFiltersResponse
	GetBody() *ListTrafficMirrorFiltersResponseBody
}

type ListTrafficMirrorFiltersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficMirrorFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficMirrorFiltersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficMirrorFiltersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficMirrorFiltersResponse) GetBody() *ListTrafficMirrorFiltersResponseBody {
	return s.Body
}

func (s *ListTrafficMirrorFiltersResponse) SetHeaders(v map[string]*string) *ListTrafficMirrorFiltersResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficMirrorFiltersResponse) SetStatusCode(v int32) *ListTrafficMirrorFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponse) SetBody(v *ListTrafficMirrorFiltersResponseBody) *ListTrafficMirrorFiltersResponse {
	s.Body = v
	return s
}

func (s *ListTrafficMirrorFiltersResponse) Validate() error {
	return dara.Validate(s)
}
