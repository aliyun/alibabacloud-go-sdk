// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOnlineEvalTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOnlineEvalTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOnlineEvalTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOnlineEvalTaskResponseBody) *UpdateOnlineEvalTaskResponse
	GetBody() *UpdateOnlineEvalTaskResponseBody
}

type UpdateOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOnlineEvalTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOnlineEvalTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOnlineEvalTaskResponse) GetBody() *UpdateOnlineEvalTaskResponseBody {
	return s.Body
}

func (s *UpdateOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *UpdateOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateOnlineEvalTaskResponse) SetStatusCode(v int32) *UpdateOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponse) SetBody(v *UpdateOnlineEvalTaskResponseBody) *UpdateOnlineEvalTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateOnlineEvalTaskResponse) Validate() error {
	return dara.Validate(s)
}
