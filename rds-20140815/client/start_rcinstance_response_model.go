// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartRCInstanceResponseBody) *StartRCInstanceResponse
	GetBody() *StartRCInstanceResponseBody
}

type StartRCInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRCInstanceResponse) GetBody() *StartRCInstanceResponseBody {
	return s.Body
}

func (s *StartRCInstanceResponse) SetHeaders(v map[string]*string) *StartRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartRCInstanceResponse) SetStatusCode(v int32) *StartRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRCInstanceResponse) SetBody(v *StartRCInstanceResponseBody) *StartRCInstanceResponse {
	s.Body = v
	return s
}

func (s *StartRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
