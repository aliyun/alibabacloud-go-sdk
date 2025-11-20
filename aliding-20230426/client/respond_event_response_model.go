// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RespondEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RespondEventResponse
	GetStatusCode() *int32
	SetBody(v *RespondEventResponseBody) *RespondEventResponse
	GetBody() *RespondEventResponseBody
}

type RespondEventResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RespondEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RespondEventResponse) String() string {
	return dara.Prettify(s)
}

func (s RespondEventResponse) GoString() string {
	return s.String()
}

func (s *RespondEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RespondEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RespondEventResponse) GetBody() *RespondEventResponseBody {
	return s.Body
}

func (s *RespondEventResponse) SetHeaders(v map[string]*string) *RespondEventResponse {
	s.Headers = v
	return s
}

func (s *RespondEventResponse) SetStatusCode(v int32) *RespondEventResponse {
	s.StatusCode = &v
	return s
}

func (s *RespondEventResponse) SetBody(v *RespondEventResponseBody) *RespondEventResponse {
	s.Body = v
	return s
}

func (s *RespondEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
