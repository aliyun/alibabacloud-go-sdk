// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *DetachEaisEiResponseBody) *DetachEaisEiResponse
	GetBody() *DetachEaisEiResponseBody
}

type DetachEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachEaisEiResponse) GoString() string {
	return s.String()
}

func (s *DetachEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachEaisEiResponse) GetBody() *DetachEaisEiResponseBody {
	return s.Body
}

func (s *DetachEaisEiResponse) SetHeaders(v map[string]*string) *DetachEaisEiResponse {
	s.Headers = v
	return s
}

func (s *DetachEaisEiResponse) SetStatusCode(v int32) *DetachEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEaisEiResponse) SetBody(v *DetachEaisEiResponseBody) *DetachEaisEiResponse {
	s.Body = v
	return s
}

func (s *DetachEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
