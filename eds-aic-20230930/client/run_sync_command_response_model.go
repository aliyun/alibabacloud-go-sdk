// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSyncCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSyncCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSyncCommandResponse
	GetStatusCode() *int32
	SetBody(v *RunSyncCommandResponseBody) *RunSyncCommandResponse
	GetBody() *RunSyncCommandResponseBody
}

type RunSyncCommandResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSyncCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSyncCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSyncCommandResponse) GoString() string {
	return s.String()
}

func (s *RunSyncCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSyncCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSyncCommandResponse) GetBody() *RunSyncCommandResponseBody {
	return s.Body
}

func (s *RunSyncCommandResponse) SetHeaders(v map[string]*string) *RunSyncCommandResponse {
	s.Headers = v
	return s
}

func (s *RunSyncCommandResponse) SetStatusCode(v int32) *RunSyncCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSyncCommandResponse) SetBody(v *RunSyncCommandResponseBody) *RunSyncCommandResponse {
	s.Body = v
	return s
}

func (s *RunSyncCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
