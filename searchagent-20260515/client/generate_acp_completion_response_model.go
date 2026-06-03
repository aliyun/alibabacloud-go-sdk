// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAcpCompletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAcpCompletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAcpCompletionResponse
	GetStatusCode() *int32
	SetId(v string) *GenerateAcpCompletionResponse
	GetId() *string
	SetEvent(v string) *GenerateAcpCompletionResponse
	GetEvent() *string
	SetBody(v string) *GenerateAcpCompletionResponse
	GetBody() *string
}

type GenerateAcpCompletionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAcpCompletionResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAcpCompletionResponse) GoString() string {
	return s.String()
}

func (s *GenerateAcpCompletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAcpCompletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAcpCompletionResponse) GetId() *string {
	return s.Id
}

func (s *GenerateAcpCompletionResponse) GetEvent() *string {
	return s.Event
}

func (s *GenerateAcpCompletionResponse) GetBody() *string {
	return s.Body
}

func (s *GenerateAcpCompletionResponse) SetHeaders(v map[string]*string) *GenerateAcpCompletionResponse {
	s.Headers = v
	return s
}

func (s *GenerateAcpCompletionResponse) SetStatusCode(v int32) *GenerateAcpCompletionResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAcpCompletionResponse) SetId(v string) *GenerateAcpCompletionResponse {
	s.Id = &v
	return s
}

func (s *GenerateAcpCompletionResponse) SetEvent(v string) *GenerateAcpCompletionResponse {
	s.Event = &v
	return s
}

func (s *GenerateAcpCompletionResponse) SetBody(v string) *GenerateAcpCompletionResponse {
	s.Body = &v
	return s
}

func (s *GenerateAcpCompletionResponse) Validate() error {
	return dara.Validate(s)
}
