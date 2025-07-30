// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEiamRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEiamRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEiamRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListEiamRegionsResponseBody) *ListEiamRegionsResponse
	GetBody() *ListEiamRegionsResponseBody
}

type ListEiamRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEiamRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEiamRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEiamRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEiamRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEiamRegionsResponse) GetBody() *ListEiamRegionsResponseBody {
	return s.Body
}

func (s *ListEiamRegionsResponse) SetHeaders(v map[string]*string) *ListEiamRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListEiamRegionsResponse) SetStatusCode(v int32) *ListEiamRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEiamRegionsResponse) SetBody(v *ListEiamRegionsResponseBody) *ListEiamRegionsResponse {
	s.Body = v
	return s
}

func (s *ListEiamRegionsResponse) Validate() error {
	return dara.Validate(s)
}
