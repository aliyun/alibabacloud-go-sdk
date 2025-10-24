// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListByURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaListByURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaListByURLResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaListByURLResponseBody) *QueryMediaListByURLResponse
	GetBody() *QueryMediaListByURLResponseBody
}

type QueryMediaListByURLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaListByURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaListByURLResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaListByURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaListByURLResponse) GetBody() *QueryMediaListByURLResponseBody {
	return s.Body
}

func (s *QueryMediaListByURLResponse) SetHeaders(v map[string]*string) *QueryMediaListByURLResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaListByURLResponse) SetStatusCode(v int32) *QueryMediaListByURLResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaListByURLResponse) SetBody(v *QueryMediaListByURLResponseBody) *QueryMediaListByURLResponse {
	s.Body = v
	return s
}

func (s *QueryMediaListByURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
