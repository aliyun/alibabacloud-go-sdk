// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutonomousNotifyEventContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutonomousNotifyEventContentResponse
	GetStatusCode() *int32
	SetBody(v *GetAutonomousNotifyEventContentResponseBody) *GetAutonomousNotifyEventContentResponse
	GetBody() *GetAutonomousNotifyEventContentResponseBody
}

type GetAutonomousNotifyEventContentResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutonomousNotifyEventContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutonomousNotifyEventContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventContentResponse) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutonomousNotifyEventContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutonomousNotifyEventContentResponse) GetBody() *GetAutonomousNotifyEventContentResponseBody {
	return s.Body
}

func (s *GetAutonomousNotifyEventContentResponse) SetHeaders(v map[string]*string) *GetAutonomousNotifyEventContentResponse {
	s.Headers = v
	return s
}

func (s *GetAutonomousNotifyEventContentResponse) SetStatusCode(v int32) *GetAutonomousNotifyEventContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponse) SetBody(v *GetAutonomousNotifyEventContentResponseBody) *GetAutonomousNotifyEventContentResponse {
	s.Body = v
	return s
}

func (s *GetAutonomousNotifyEventContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
