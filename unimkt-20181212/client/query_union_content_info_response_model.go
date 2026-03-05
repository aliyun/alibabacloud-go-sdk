// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionContentInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUnionContentInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUnionContentInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryUnionContentInfoResponseBody) *QueryUnionContentInfoResponse
	GetBody() *QueryUnionContentInfoResponseBody
}

type QueryUnionContentInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnionContentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnionContentInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionContentInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUnionContentInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUnionContentInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUnionContentInfoResponse) GetBody() *QueryUnionContentInfoResponseBody {
	return s.Body
}

func (s *QueryUnionContentInfoResponse) SetHeaders(v map[string]*string) *QueryUnionContentInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUnionContentInfoResponse) SetStatusCode(v int32) *QueryUnionContentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnionContentInfoResponse) SetBody(v *QueryUnionContentInfoResponseBody) *QueryUnionContentInfoResponse {
	s.Body = v
	return s
}

func (s *QueryUnionContentInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
