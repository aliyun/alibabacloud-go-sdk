// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIDCImportCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIDCImportCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIDCImportCommandResponse
	GetStatusCode() *int32
	SetBody(v *CreateIDCImportCommandResponseBody) *CreateIDCImportCommandResponse
	GetBody() *CreateIDCImportCommandResponseBody
}

type CreateIDCImportCommandResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIDCImportCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIDCImportCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIDCImportCommandResponse) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIDCImportCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIDCImportCommandResponse) GetBody() *CreateIDCImportCommandResponseBody {
	return s.Body
}

func (s *CreateIDCImportCommandResponse) SetHeaders(v map[string]*string) *CreateIDCImportCommandResponse {
	s.Headers = v
	return s
}

func (s *CreateIDCImportCommandResponse) SetStatusCode(v int32) *CreateIDCImportCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIDCImportCommandResponse) SetBody(v *CreateIDCImportCommandResponseBody) *CreateIDCImportCommandResponse {
	s.Body = v
	return s
}

func (s *CreateIDCImportCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
