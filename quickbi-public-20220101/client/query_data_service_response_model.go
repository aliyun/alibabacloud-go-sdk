// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataServiceResponseBody) *QueryDataServiceResponse
	GetBody() *QueryDataServiceResponseBody
}

type QueryDataServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataServiceResponse) GetBody() *QueryDataServiceResponseBody {
	return s.Body
}

func (s *QueryDataServiceResponse) SetHeaders(v map[string]*string) *QueryDataServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryDataServiceResponse) SetStatusCode(v int32) *QueryDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataServiceResponse) SetBody(v *QueryDataServiceResponseBody) *QueryDataServiceResponse {
	s.Body = v
	return s
}

func (s *QueryDataServiceResponse) Validate() error {
	return dara.Validate(s)
}
