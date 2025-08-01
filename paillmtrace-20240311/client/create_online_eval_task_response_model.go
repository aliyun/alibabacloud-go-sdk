// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineEvalTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOnlineEvalTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOnlineEvalTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOnlineEvalTaskResponseBody) *CreateOnlineEvalTaskResponse
	GetBody() *CreateOnlineEvalTaskResponseBody
}

type CreateOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineEvalTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOnlineEvalTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOnlineEvalTaskResponse) GetBody() *CreateOnlineEvalTaskResponseBody {
	return s.Body
}

func (s *CreateOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *CreateOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOnlineEvalTaskResponse) SetStatusCode(v int32) *CreateOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOnlineEvalTaskResponse) SetBody(v *CreateOnlineEvalTaskResponseBody) *CreateOnlineEvalTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOnlineEvalTaskResponse) Validate() error {
	return dara.Validate(s)
}
