// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDBProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDBProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDBProjectResponseBody) *DeleteAgenticDBProjectResponse
	GetBody() *DeleteAgenticDBProjectResponseBody
}

type DeleteAgenticDBProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDBProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDBProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDBProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDBProjectResponse) GetBody() *DeleteAgenticDBProjectResponseBody {
	return s.Body
}

func (s *DeleteAgenticDBProjectResponse) SetHeaders(v map[string]*string) *DeleteAgenticDBProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDBProjectResponse) SetStatusCode(v int32) *DeleteAgenticDBProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDBProjectResponse) SetBody(v *DeleteAgenticDBProjectResponseBody) *DeleteAgenticDBProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDBProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
