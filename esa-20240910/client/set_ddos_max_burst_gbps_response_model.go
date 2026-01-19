// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDdosMaxBurstGbpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDdosMaxBurstGbpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDdosMaxBurstGbpsResponse
	GetStatusCode() *int32
	SetBody(v *SetDdosMaxBurstGbpsResponseBody) *SetDdosMaxBurstGbpsResponse
	GetBody() *SetDdosMaxBurstGbpsResponseBody
}

type SetDdosMaxBurstGbpsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDdosMaxBurstGbpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDdosMaxBurstGbpsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDdosMaxBurstGbpsResponse) GoString() string {
	return s.String()
}

func (s *SetDdosMaxBurstGbpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDdosMaxBurstGbpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDdosMaxBurstGbpsResponse) GetBody() *SetDdosMaxBurstGbpsResponseBody {
	return s.Body
}

func (s *SetDdosMaxBurstGbpsResponse) SetHeaders(v map[string]*string) *SetDdosMaxBurstGbpsResponse {
	s.Headers = v
	return s
}

func (s *SetDdosMaxBurstGbpsResponse) SetStatusCode(v int32) *SetDdosMaxBurstGbpsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDdosMaxBurstGbpsResponse) SetBody(v *SetDdosMaxBurstGbpsResponseBody) *SetDdosMaxBurstGbpsResponse {
	s.Body = v
	return s
}

func (s *SetDdosMaxBurstGbpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
