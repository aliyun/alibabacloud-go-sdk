// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByOuterIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBindsByOuterIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBindsByOuterIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryBindsByOuterIdResponseBody) *QueryBindsByOuterIdResponse
	GetBody() *QueryBindsByOuterIdResponseBody
}

type QueryBindsByOuterIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBindsByOuterIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBindsByOuterIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByOuterIdResponse) GoString() string {
	return s.String()
}

func (s *QueryBindsByOuterIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBindsByOuterIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBindsByOuterIdResponse) GetBody() *QueryBindsByOuterIdResponseBody {
	return s.Body
}

func (s *QueryBindsByOuterIdResponse) SetHeaders(v map[string]*string) *QueryBindsByOuterIdResponse {
	s.Headers = v
	return s
}

func (s *QueryBindsByOuterIdResponse) SetStatusCode(v int32) *QueryBindsByOuterIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBindsByOuterIdResponse) SetBody(v *QueryBindsByOuterIdResponseBody) *QueryBindsByOuterIdResponse {
	s.Body = v
	return s
}

func (s *QueryBindsByOuterIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
