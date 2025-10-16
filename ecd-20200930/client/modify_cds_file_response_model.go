// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdsFileResponseBody) *ModifyCdsFileResponse
	GetBody() *ModifyCdsFileResponseBody
}

type ModifyCdsFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdsFileResponse) GetBody() *ModifyCdsFileResponseBody {
	return s.Body
}

func (s *ModifyCdsFileResponse) SetHeaders(v map[string]*string) *ModifyCdsFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdsFileResponse) SetStatusCode(v int32) *ModifyCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdsFileResponse) SetBody(v *ModifyCdsFileResponseBody) *ModifyCdsFileResponse {
	s.Body = v
	return s
}

func (s *ModifyCdsFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
