// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedeemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRedeemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRedeemResponse
	GetStatusCode() *int32
	SetBody(v *QueryRedeemResponseBody) *QueryRedeemResponse
	GetBody() *QueryRedeemResponseBody
}

type QueryRedeemResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRedeemResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemResponse) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRedeemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRedeemResponse) GetBody() *QueryRedeemResponseBody {
	return s.Body
}

func (s *QueryRedeemResponse) SetHeaders(v map[string]*string) *QueryRedeemResponse {
	s.Headers = v
	return s
}

func (s *QueryRedeemResponse) SetStatusCode(v int32) *QueryRedeemResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRedeemResponse) SetBody(v *QueryRedeemResponseBody) *QueryRedeemResponse {
	s.Body = v
	return s
}

func (s *QueryRedeemResponse) Validate() error {
	return dara.Validate(s)
}
