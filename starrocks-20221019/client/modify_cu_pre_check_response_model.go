// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCuPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCuPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCuPreCheckResponseBody) *ModifyCuPreCheckResponse
	GetBody() *ModifyCuPreCheckResponseBody
}

type ModifyCuPreCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCuPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCuPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuPreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCuPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCuPreCheckResponse) GetBody() *ModifyCuPreCheckResponseBody {
	return s.Body
}

func (s *ModifyCuPreCheckResponse) SetHeaders(v map[string]*string) *ModifyCuPreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyCuPreCheckResponse) SetStatusCode(v int32) *ModifyCuPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCuPreCheckResponse) SetBody(v *ModifyCuPreCheckResponseBody) *ModifyCuPreCheckResponse {
	s.Body = v
	return s
}

func (s *ModifyCuPreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
