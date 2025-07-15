// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachCenResponse
	GetStatusCode() *int32
	SetBody(v *DetachCenResponseBody) *DetachCenResponse
	GetBody() *DetachCenResponseBody
}

type DetachCenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachCenResponse) GoString() string {
	return s.String()
}

func (s *DetachCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachCenResponse) GetBody() *DetachCenResponseBody {
	return s.Body
}

func (s *DetachCenResponse) SetHeaders(v map[string]*string) *DetachCenResponse {
	s.Headers = v
	return s
}

func (s *DetachCenResponse) SetStatusCode(v int32) *DetachCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachCenResponse) SetBody(v *DetachCenResponseBody) *DetachCenResponse {
	s.Body = v
	return s
}

func (s *DetachCenResponse) Validate() error {
	return dara.Validate(s)
}
