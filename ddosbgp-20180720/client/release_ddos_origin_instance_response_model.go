// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDdosOriginInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseDdosOriginInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseDdosOriginInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseDdosOriginInstanceResponseBody) *ReleaseDdosOriginInstanceResponse
	GetBody() *ReleaseDdosOriginInstanceResponseBody
}

type ReleaseDdosOriginInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseDdosOriginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseDdosOriginInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDdosOriginInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseDdosOriginInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseDdosOriginInstanceResponse) GetBody() *ReleaseDdosOriginInstanceResponseBody {
	return s.Body
}

func (s *ReleaseDdosOriginInstanceResponse) SetHeaders(v map[string]*string) *ReleaseDdosOriginInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseDdosOriginInstanceResponse) SetStatusCode(v int32) *ReleaseDdosOriginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseDdosOriginInstanceResponse) SetBody(v *ReleaseDdosOriginInstanceResponseBody) *ReleaseDdosOriginInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseDdosOriginInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
