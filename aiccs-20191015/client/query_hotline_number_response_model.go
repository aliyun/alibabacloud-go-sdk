// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHotlineNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHotlineNumberResponse
	GetStatusCode() *int32
	SetBody(v *QueryHotlineNumberResponseBody) *QueryHotlineNumberResponse
	GetBody() *QueryHotlineNumberResponseBody
}

type QueryHotlineNumberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotlineNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHotlineNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHotlineNumberResponse) GetBody() *QueryHotlineNumberResponseBody {
	return s.Body
}

func (s *QueryHotlineNumberResponse) SetHeaders(v map[string]*string) *QueryHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineNumberResponse) SetStatusCode(v int32) *QueryHotlineNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotlineNumberResponse) SetBody(v *QueryHotlineNumberResponseBody) *QueryHotlineNumberResponse {
	s.Body = v
	return s
}

func (s *QueryHotlineNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
