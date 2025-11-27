// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StartRCInstancesResponseBody) *StartRCInstancesResponse
	GetBody() *StartRCInstancesResponseBody
}

type StartRCInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *StartRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRCInstancesResponse) GetBody() *StartRCInstancesResponseBody {
	return s.Body
}

func (s *StartRCInstancesResponse) SetHeaders(v map[string]*string) *StartRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *StartRCInstancesResponse) SetStatusCode(v int32) *StartRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRCInstancesResponse) SetBody(v *StartRCInstancesResponseBody) *StartRCInstancesResponse {
	s.Body = v
	return s
}

func (s *StartRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
