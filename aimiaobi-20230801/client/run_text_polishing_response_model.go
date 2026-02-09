// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTextPolishingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTextPolishingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTextPolishingResponse
	GetStatusCode() *int32
	SetId(v string) *RunTextPolishingResponse
	GetId() *string
	SetEvent(v string) *RunTextPolishingResponse
	GetEvent() *string
	SetBody(v *RunTextPolishingResponseBody) *RunTextPolishingResponse
	GetBody() *RunTextPolishingResponseBody
}

type RunTextPolishingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunTextPolishingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTextPolishingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponse) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTextPolishingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTextPolishingResponse) GetId() *string {
	return s.Id
}

func (s *RunTextPolishingResponse) GetEvent() *string {
	return s.Event
}

func (s *RunTextPolishingResponse) GetBody() *RunTextPolishingResponseBody {
	return s.Body
}

func (s *RunTextPolishingResponse) SetHeaders(v map[string]*string) *RunTextPolishingResponse {
	s.Headers = v
	return s
}

func (s *RunTextPolishingResponse) SetStatusCode(v int32) *RunTextPolishingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTextPolishingResponse) SetId(v string) *RunTextPolishingResponse {
	s.Id = &v
	return s
}

func (s *RunTextPolishingResponse) SetEvent(v string) *RunTextPolishingResponse {
	s.Event = &v
	return s
}

func (s *RunTextPolishingResponse) SetBody(v *RunTextPolishingResponseBody) *RunTextPolishingResponse {
	s.Body = v
	return s
}

func (s *RunTextPolishingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
