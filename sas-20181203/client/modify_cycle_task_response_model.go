// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCycleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCycleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCycleTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCycleTaskResponseBody) *ModifyCycleTaskResponse
	GetBody() *ModifyCycleTaskResponseBody
}

type ModifyCycleTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCycleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCycleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCycleTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyCycleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCycleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCycleTaskResponse) GetBody() *ModifyCycleTaskResponseBody {
	return s.Body
}

func (s *ModifyCycleTaskResponse) SetHeaders(v map[string]*string) *ModifyCycleTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyCycleTaskResponse) SetStatusCode(v int32) *ModifyCycleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCycleTaskResponse) SetBody(v *ModifyCycleTaskResponseBody) *ModifyCycleTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyCycleTaskResponse) Validate() error {
	return dara.Validate(s)
}
