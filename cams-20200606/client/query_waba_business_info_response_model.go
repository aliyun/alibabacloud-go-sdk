// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWabaBusinessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWabaBusinessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWabaBusinessInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryWabaBusinessInfoResponseBody) *QueryWabaBusinessInfoResponse
	GetBody() *QueryWabaBusinessInfoResponseBody
}

type QueryWabaBusinessInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWabaBusinessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWabaBusinessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWabaBusinessInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWabaBusinessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWabaBusinessInfoResponse) GetBody() *QueryWabaBusinessInfoResponseBody {
	return s.Body
}

func (s *QueryWabaBusinessInfoResponse) SetHeaders(v map[string]*string) *QueryWabaBusinessInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryWabaBusinessInfoResponse) SetStatusCode(v int32) *QueryWabaBusinessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWabaBusinessInfoResponse) SetBody(v *QueryWabaBusinessInfoResponseBody) *QueryWabaBusinessInfoResponse {
	s.Body = v
	return s
}

func (s *QueryWabaBusinessInfoResponse) Validate() error {
	return dara.Validate(s)
}
