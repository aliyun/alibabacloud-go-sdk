// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceResponseBody) *QueryInstanceResponse
	GetBody() *QueryInstanceResponseBody
}

type QueryInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceResponse) GetBody() *QueryInstanceResponseBody {
	return s.Body
}

func (s *QueryInstanceResponse) SetHeaders(v map[string]*string) *QueryInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceResponse) SetStatusCode(v int32) *QueryInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceResponse) SetBody(v *QueryInstanceResponseBody) *QueryInstanceResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
