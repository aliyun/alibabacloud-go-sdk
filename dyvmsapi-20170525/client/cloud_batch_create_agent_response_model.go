// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchCreateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudBatchCreateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudBatchCreateAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudBatchCreateAgentResponseBody) *CloudBatchCreateAgentResponse
	GetBody() *CloudBatchCreateAgentResponseBody
}

type CloudBatchCreateAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudBatchCreateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudBatchCreateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchCreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudBatchCreateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudBatchCreateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudBatchCreateAgentResponse) GetBody() *CloudBatchCreateAgentResponseBody {
	return s.Body
}

func (s *CloudBatchCreateAgentResponse) SetHeaders(v map[string]*string) *CloudBatchCreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudBatchCreateAgentResponse) SetStatusCode(v int32) *CloudBatchCreateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudBatchCreateAgentResponse) SetBody(v *CloudBatchCreateAgentResponseBody) *CloudBatchCreateAgentResponse {
	s.Body = v
	return s
}

func (s *CloudBatchCreateAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
