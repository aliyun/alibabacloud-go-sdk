// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommercializeFetchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommercializeFetchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommercializeFetchResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *CommercializeFetchResponse
	GetBody() map[string]interface{}
}

type CommercializeFetchResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommercializeFetchResponse) String() string {
	return dara.Prettify(s)
}

func (s CommercializeFetchResponse) GoString() string {
	return s.String()
}

func (s *CommercializeFetchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommercializeFetchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommercializeFetchResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *CommercializeFetchResponse) SetHeaders(v map[string]*string) *CommercializeFetchResponse {
	s.Headers = v
	return s
}

func (s *CommercializeFetchResponse) SetStatusCode(v int32) *CommercializeFetchResponse {
	s.StatusCode = &v
	return s
}

func (s *CommercializeFetchResponse) SetBody(v map[string]interface{}) *CommercializeFetchResponse {
	s.Body = v
	return s
}

func (s *CommercializeFetchResponse) Validate() error {
	return dara.Validate(s)
}
