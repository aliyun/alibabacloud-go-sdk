// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterSpecificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryClusterSpecificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryClusterSpecificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryClusterSpecificationResponseBody) *QueryClusterSpecificationResponse
	GetBody() *QueryClusterSpecificationResponseBody
}

type QueryClusterSpecificationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryClusterSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryClusterSpecificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterSpecificationResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryClusterSpecificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryClusterSpecificationResponse) GetBody() *QueryClusterSpecificationResponseBody {
	return s.Body
}

func (s *QueryClusterSpecificationResponse) SetHeaders(v map[string]*string) *QueryClusterSpecificationResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterSpecificationResponse) SetStatusCode(v int32) *QueryClusterSpecificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClusterSpecificationResponse) SetBody(v *QueryClusterSpecificationResponseBody) *QueryClusterSpecificationResponse {
	s.Body = v
	return s
}

func (s *QueryClusterSpecificationResponse) Validate() error {
	return dara.Validate(s)
}
