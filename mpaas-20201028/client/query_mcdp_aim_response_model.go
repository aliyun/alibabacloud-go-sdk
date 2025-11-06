// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpAimResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcdpAimResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcdpAimResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcdpAimResponseBody) *QueryMcdpAimResponse
	GetBody() *QueryMcdpAimResponseBody
}

type QueryMcdpAimResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcdpAimResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcdpAimResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpAimResponse) GoString() string {
	return s.String()
}

func (s *QueryMcdpAimResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcdpAimResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcdpAimResponse) GetBody() *QueryMcdpAimResponseBody {
	return s.Body
}

func (s *QueryMcdpAimResponse) SetHeaders(v map[string]*string) *QueryMcdpAimResponse {
	s.Headers = v
	return s
}

func (s *QueryMcdpAimResponse) SetStatusCode(v int32) *QueryMcdpAimResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcdpAimResponse) SetBody(v *QueryMcdpAimResponseBody) *QueryMcdpAimResponse {
	s.Body = v
	return s
}

func (s *QueryMcdpAimResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
