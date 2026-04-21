// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OmniAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OmniAnswerResponse
	GetStatusCode() *int32
	SetId(v string) *OmniAnswerResponse
	GetId() *string
	SetEvent(v string) *OmniAnswerResponse
	GetEvent() *string
	SetBody(v string) *OmniAnswerResponse
	GetBody() *string
}

type OmniAnswerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OmniAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s OmniAnswerResponse) GoString() string {
	return s.String()
}

func (s *OmniAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OmniAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OmniAnswerResponse) GetId() *string {
	return s.Id
}

func (s *OmniAnswerResponse) GetEvent() *string {
	return s.Event
}

func (s *OmniAnswerResponse) GetBody() *string {
	return s.Body
}

func (s *OmniAnswerResponse) SetHeaders(v map[string]*string) *OmniAnswerResponse {
	s.Headers = v
	return s
}

func (s *OmniAnswerResponse) SetStatusCode(v int32) *OmniAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *OmniAnswerResponse) SetId(v string) *OmniAnswerResponse {
	s.Id = &v
	return s
}

func (s *OmniAnswerResponse) SetEvent(v string) *OmniAnswerResponse {
	s.Event = &v
	return s
}

func (s *OmniAnswerResponse) SetBody(v string) *OmniAnswerResponse {
	s.Body = &v
	return s
}

func (s *OmniAnswerResponse) Validate() error {
	return dara.Validate(s)
}
