// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationAccessIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryApplicationAccessIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryApplicationAccessIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryApplicationAccessIdResponseBody) *QueryApplicationAccessIdResponse
	GetBody() *QueryApplicationAccessIdResponseBody
}

type QueryApplicationAccessIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApplicationAccessIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApplicationAccessIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationAccessIdResponse) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryApplicationAccessIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryApplicationAccessIdResponse) GetBody() *QueryApplicationAccessIdResponseBody {
	return s.Body
}

func (s *QueryApplicationAccessIdResponse) SetHeaders(v map[string]*string) *QueryApplicationAccessIdResponse {
	s.Headers = v
	return s
}

func (s *QueryApplicationAccessIdResponse) SetStatusCode(v int32) *QueryApplicationAccessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponse) SetBody(v *QueryApplicationAccessIdResponseBody) *QueryApplicationAccessIdResponse {
	s.Body = v
	return s
}

func (s *QueryApplicationAccessIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
