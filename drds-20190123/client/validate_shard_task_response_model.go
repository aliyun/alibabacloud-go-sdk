// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShardTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateShardTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateShardTaskResponse
	GetStatusCode() *int32
	SetBody(v *ValidateShardTaskResponseBody) *ValidateShardTaskResponse
	GetBody() *ValidateShardTaskResponseBody
}

type ValidateShardTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateShardTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateShardTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateShardTaskResponse) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateShardTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateShardTaskResponse) GetBody() *ValidateShardTaskResponseBody {
	return s.Body
}

func (s *ValidateShardTaskResponse) SetHeaders(v map[string]*string) *ValidateShardTaskResponse {
	s.Headers = v
	return s
}

func (s *ValidateShardTaskResponse) SetStatusCode(v int32) *ValidateShardTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateShardTaskResponse) SetBody(v *ValidateShardTaskResponseBody) *ValidateShardTaskResponse {
	s.Body = v
	return s
}

func (s *ValidateShardTaskResponse) Validate() error {
	return dara.Validate(s)
}
