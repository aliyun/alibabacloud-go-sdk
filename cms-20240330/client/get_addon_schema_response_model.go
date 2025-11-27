// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAddonSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAddonSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetAddonSchemaResponseBody) *GetAddonSchemaResponse
	GetBody() *GetAddonSchemaResponseBody
}

type GetAddonSchemaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddonSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddonSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAddonSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAddonSchemaResponse) GetBody() *GetAddonSchemaResponseBody {
	return s.Body
}

func (s *GetAddonSchemaResponse) SetHeaders(v map[string]*string) *GetAddonSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetAddonSchemaResponse) SetStatusCode(v int32) *GetAddonSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddonSchemaResponse) SetBody(v *GetAddonSchemaResponseBody) *GetAddonSchemaResponse {
	s.Body = v
	return s
}

func (s *GetAddonSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
