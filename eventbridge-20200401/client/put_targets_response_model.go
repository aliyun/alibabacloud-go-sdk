// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutTargetsResponse
	GetStatusCode() *int32
	SetBody(v *PutTargetsResponseBody) *PutTargetsResponse
	GetBody() *PutTargetsResponseBody
}

type PutTargetsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutTargetsResponse) GetBody() *PutTargetsResponseBody {
	return s.Body
}

func (s *PutTargetsResponse) SetHeaders(v map[string]*string) *PutTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutTargetsResponse) SetStatusCode(v int32) *PutTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTargetsResponse) SetBody(v *PutTargetsResponseBody) *PutTargetsResponse {
	s.Body = v
	return s
}

func (s *PutTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
