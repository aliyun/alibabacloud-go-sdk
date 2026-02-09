// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAbbreviationContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunAbbreviationContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunAbbreviationContentResponse
	GetStatusCode() *int32
	SetId(v string) *RunAbbreviationContentResponse
	GetId() *string
	SetEvent(v string) *RunAbbreviationContentResponse
	GetEvent() *string
	SetBody(v *RunAbbreviationContentResponseBody) *RunAbbreviationContentResponse
	GetBody() *RunAbbreviationContentResponseBody
}

type RunAbbreviationContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunAbbreviationContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAbbreviationContentResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponse) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunAbbreviationContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAbbreviationContentResponse) GetId() *string {
	return s.Id
}

func (s *RunAbbreviationContentResponse) GetEvent() *string {
	return s.Event
}

func (s *RunAbbreviationContentResponse) GetBody() *RunAbbreviationContentResponseBody {
	return s.Body
}

func (s *RunAbbreviationContentResponse) SetHeaders(v map[string]*string) *RunAbbreviationContentResponse {
	s.Headers = v
	return s
}

func (s *RunAbbreviationContentResponse) SetStatusCode(v int32) *RunAbbreviationContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAbbreviationContentResponse) SetId(v string) *RunAbbreviationContentResponse {
	s.Id = &v
	return s
}

func (s *RunAbbreviationContentResponse) SetEvent(v string) *RunAbbreviationContentResponse {
	s.Event = &v
	return s
}

func (s *RunAbbreviationContentResponse) SetBody(v *RunAbbreviationContentResponseBody) *RunAbbreviationContentResponse {
	s.Body = v
	return s
}

func (s *RunAbbreviationContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
