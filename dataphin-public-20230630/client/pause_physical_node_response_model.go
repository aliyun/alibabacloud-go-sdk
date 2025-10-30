// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePhysicalNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PausePhysicalNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PausePhysicalNodeResponse
	GetStatusCode() *int32
	SetBody(v *PausePhysicalNodeResponseBody) *PausePhysicalNodeResponse
	GetBody() *PausePhysicalNodeResponseBody
}

type PausePhysicalNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PausePhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PausePhysicalNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PausePhysicalNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PausePhysicalNodeResponse) GetBody() *PausePhysicalNodeResponseBody {
	return s.Body
}

func (s *PausePhysicalNodeResponse) SetHeaders(v map[string]*string) *PausePhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *PausePhysicalNodeResponse) SetStatusCode(v int32) *PausePhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *PausePhysicalNodeResponse) SetBody(v *PausePhysicalNodeResponseBody) *PausePhysicalNodeResponse {
	s.Body = v
	return s
}

func (s *PausePhysicalNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
