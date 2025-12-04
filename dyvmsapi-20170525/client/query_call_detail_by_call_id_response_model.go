// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByCallIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallDetailByCallIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallDetailByCallIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallDetailByCallIdResponseBody) *QueryCallDetailByCallIdResponse
	GetBody() *QueryCallDetailByCallIdResponseBody
}

type QueryCallDetailByCallIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallDetailByCallIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallDetailByCallIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByCallIdResponse) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallDetailByCallIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallDetailByCallIdResponse) GetBody() *QueryCallDetailByCallIdResponseBody {
	return s.Body
}

func (s *QueryCallDetailByCallIdResponse) SetHeaders(v map[string]*string) *QueryCallDetailByCallIdResponse {
	s.Headers = v
	return s
}

func (s *QueryCallDetailByCallIdResponse) SetStatusCode(v int32) *QueryCallDetailByCallIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallDetailByCallIdResponse) SetBody(v *QueryCallDetailByCallIdResponseBody) *QueryCallDetailByCallIdResponse {
	s.Body = v
	return s
}

func (s *QueryCallDetailByCallIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
