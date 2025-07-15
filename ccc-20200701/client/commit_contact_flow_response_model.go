// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitContactFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitContactFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitContactFlowResponse
	GetStatusCode() *int32
	SetBody(v *CommitContactFlowResponseBody) *CommitContactFlowResponse
	GetBody() *CommitContactFlowResponseBody
}

type CommitContactFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitContactFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitContactFlowResponse) GoString() string {
	return s.String()
}

func (s *CommitContactFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitContactFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitContactFlowResponse) GetBody() *CommitContactFlowResponseBody {
	return s.Body
}

func (s *CommitContactFlowResponse) SetHeaders(v map[string]*string) *CommitContactFlowResponse {
	s.Headers = v
	return s
}

func (s *CommitContactFlowResponse) SetStatusCode(v int32) *CommitContactFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitContactFlowResponse) SetBody(v *CommitContactFlowResponseBody) *CommitContactFlowResponse {
	s.Body = v
	return s
}

func (s *CommitContactFlowResponse) Validate() error {
	return dara.Validate(s)
}
