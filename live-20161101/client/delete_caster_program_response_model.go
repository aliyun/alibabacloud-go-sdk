// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterProgramResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterProgramResponseBody) *DeleteCasterProgramResponse
	GetBody() *DeleteCasterProgramResponseBody
}

type DeleteCasterProgramResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterProgramResponse) GetBody() *DeleteCasterProgramResponseBody {
	return s.Body
}

func (s *DeleteCasterProgramResponse) SetHeaders(v map[string]*string) *DeleteCasterProgramResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterProgramResponse) SetStatusCode(v int32) *DeleteCasterProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterProgramResponse) SetBody(v *DeleteCasterProgramResponseBody) *DeleteCasterProgramResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterProgramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
