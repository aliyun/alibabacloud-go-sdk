// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListCollectionsResponseBody) *ListCollectionsResponse
	GetBody() *ListCollectionsResponseBody
}

type ListCollectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCollectionsResponse) GetBody() *ListCollectionsResponseBody {
	return s.Body
}

func (s *ListCollectionsResponse) SetHeaders(v map[string]*string) *ListCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListCollectionsResponse) SetStatusCode(v int32) *ListCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollectionsResponse) SetBody(v *ListCollectionsResponseBody) *ListCollectionsResponse {
	s.Body = v
	return s
}

func (s *ListCollectionsResponse) Validate() error {
	return dara.Validate(s)
}
