// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountInfoQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TBAccountInfoQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TBAccountInfoQueryResponse
	GetStatusCode() *int32
	SetBody(v *TBAccountInfoQueryResponseBody) *TBAccountInfoQueryResponse
	GetBody() *TBAccountInfoQueryResponseBody
}

type TBAccountInfoQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TBAccountInfoQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TBAccountInfoQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TBAccountInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *TBAccountInfoQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TBAccountInfoQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TBAccountInfoQueryResponse) GetBody() *TBAccountInfoQueryResponseBody {
	return s.Body
}

func (s *TBAccountInfoQueryResponse) SetHeaders(v map[string]*string) *TBAccountInfoQueryResponse {
	s.Headers = v
	return s
}

func (s *TBAccountInfoQueryResponse) SetStatusCode(v int32) *TBAccountInfoQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TBAccountInfoQueryResponse) SetBody(v *TBAccountInfoQueryResponseBody) *TBAccountInfoQueryResponse {
	s.Body = v
	return s
}

func (s *TBAccountInfoQueryResponse) Validate() error {
	return dara.Validate(s)
}
