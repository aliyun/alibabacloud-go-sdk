// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentSelectiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAndPublishAgentSelectiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAndPublishAgentSelectiveResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAndPublishAgentSelectiveResponseBody) *UpdateAndPublishAgentSelectiveResponse
	GetBody() *UpdateAndPublishAgentSelectiveResponseBody
}

type UpdateAndPublishAgentSelectiveResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAndPublishAgentSelectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAndPublishAgentSelectiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAndPublishAgentSelectiveResponse) GetBody() *UpdateAndPublishAgentSelectiveResponseBody {
	return s.Body
}

func (s *UpdateAndPublishAgentSelectiveResponse) SetHeaders(v map[string]*string) *UpdateAndPublishAgentSelectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponse) SetStatusCode(v int32) *UpdateAndPublishAgentSelectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponse) SetBody(v *UpdateAndPublishAgentSelectiveResponseBody) *UpdateAndPublishAgentSelectiveResponse {
	s.Body = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponse) Validate() error {
	return dara.Validate(s)
}
