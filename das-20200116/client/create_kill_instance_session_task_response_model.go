// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKillInstanceSessionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKillInstanceSessionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateKillInstanceSessionTaskResponseBody) *CreateKillInstanceSessionTaskResponse
	GetBody() *CreateKillInstanceSessionTaskResponseBody
}

type CreateKillInstanceSessionTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKillInstanceSessionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKillInstanceSessionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKillInstanceSessionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKillInstanceSessionTaskResponse) GetBody() *CreateKillInstanceSessionTaskResponseBody {
	return s.Body
}

func (s *CreateKillInstanceSessionTaskResponse) SetHeaders(v map[string]*string) *CreateKillInstanceSessionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateKillInstanceSessionTaskResponse) SetStatusCode(v int32) *CreateKillInstanceSessionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKillInstanceSessionTaskResponse) SetBody(v *CreateKillInstanceSessionTaskResponseBody) *CreateKillInstanceSessionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateKillInstanceSessionTaskResponse) Validate() error {
	return dara.Validate(s)
}
