// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateShardTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateShardTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateShardTaskResponseBody) *CreateShardTaskResponse
	GetBody() *CreateShardTaskResponseBody
}

type CreateShardTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShardTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateShardTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateShardTaskResponse) GetBody() *CreateShardTaskResponseBody {
	return s.Body
}

func (s *CreateShardTaskResponse) SetHeaders(v map[string]*string) *CreateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateShardTaskResponse) SetStatusCode(v int32) *CreateShardTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShardTaskResponse) SetBody(v *CreateShardTaskResponseBody) *CreateShardTaskResponse {
	s.Body = v
	return s
}

func (s *CreateShardTaskResponse) Validate() error {
	return dara.Validate(s)
}
