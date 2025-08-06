// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBagRemainingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBagRemainingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBagRemainingResponse
	GetStatusCode() *int32
	SetBody(v *QueryBagRemainingResponseBody) *QueryBagRemainingResponse
	GetBody() *QueryBagRemainingResponseBody
}

type QueryBagRemainingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBagRemainingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBagRemainingResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBagRemainingResponse) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBagRemainingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBagRemainingResponse) GetBody() *QueryBagRemainingResponseBody {
	return s.Body
}

func (s *QueryBagRemainingResponse) SetHeaders(v map[string]*string) *QueryBagRemainingResponse {
	s.Headers = v
	return s
}

func (s *QueryBagRemainingResponse) SetStatusCode(v int32) *QueryBagRemainingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBagRemainingResponse) SetBody(v *QueryBagRemainingResponseBody) *QueryBagRemainingResponse {
	s.Body = v
	return s
}

func (s *QueryBagRemainingResponse) Validate() error {
	return dara.Validate(s)
}
