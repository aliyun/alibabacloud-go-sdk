// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepoTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepoTriggerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepoTriggerResponseBody) *DeleteRepoTriggerResponse
	GetBody() *DeleteRepoTriggerResponseBody
}

type DeleteRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepoTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepoTriggerResponse) GetBody() *DeleteRepoTriggerResponseBody {
	return s.Body
}

func (s *DeleteRepoTriggerResponse) SetHeaders(v map[string]*string) *DeleteRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoTriggerResponse) SetStatusCode(v int32) *DeleteRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoTriggerResponse) SetBody(v *DeleteRepoTriggerResponseBody) *DeleteRepoTriggerResponse {
	s.Body = v
	return s
}

func (s *DeleteRepoTriggerResponse) Validate() error {
	return dara.Validate(s)
}
