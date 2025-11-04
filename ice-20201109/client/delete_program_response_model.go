// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProgramResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProgramResponseBody) *DeleteProgramResponse
	GetBody() *DeleteProgramResponseBody
}

type DeleteProgramResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProgramResponse) GoString() string {
	return s.String()
}

func (s *DeleteProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProgramResponse) GetBody() *DeleteProgramResponseBody {
	return s.Body
}

func (s *DeleteProgramResponse) SetHeaders(v map[string]*string) *DeleteProgramResponse {
	s.Headers = v
	return s
}

func (s *DeleteProgramResponse) SetStatusCode(v int32) *DeleteProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProgramResponse) SetBody(v *DeleteProgramResponseBody) *DeleteProgramResponse {
	s.Body = v
	return s
}

func (s *DeleteProgramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
