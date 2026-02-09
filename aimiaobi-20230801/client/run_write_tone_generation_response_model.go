// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWriteToneGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunWriteToneGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunWriteToneGenerationResponse
	GetStatusCode() *int32
	SetId(v string) *RunWriteToneGenerationResponse
	GetId() *string
	SetEvent(v string) *RunWriteToneGenerationResponse
	GetEvent() *string
	SetBody(v *RunWriteToneGenerationResponseBody) *RunWriteToneGenerationResponse
	GetBody() *RunWriteToneGenerationResponseBody
}

type RunWriteToneGenerationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunWriteToneGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWriteToneGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunWriteToneGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWriteToneGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunWriteToneGenerationResponse) GetEvent() *string {
	return s.Event
}

func (s *RunWriteToneGenerationResponse) GetBody() *RunWriteToneGenerationResponseBody {
	return s.Body
}

func (s *RunWriteToneGenerationResponse) SetHeaders(v map[string]*string) *RunWriteToneGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunWriteToneGenerationResponse) SetStatusCode(v int32) *RunWriteToneGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWriteToneGenerationResponse) SetId(v string) *RunWriteToneGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunWriteToneGenerationResponse) SetEvent(v string) *RunWriteToneGenerationResponse {
	s.Event = &v
	return s
}

func (s *RunWriteToneGenerationResponse) SetBody(v *RunWriteToneGenerationResponseBody) *RunWriteToneGenerationResponse {
	s.Body = v
	return s
}

func (s *RunWriteToneGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
