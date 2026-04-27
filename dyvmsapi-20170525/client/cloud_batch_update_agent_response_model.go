// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchUpdateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudBatchUpdateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudBatchUpdateAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudBatchUpdateAgentResponseBody) *CloudBatchUpdateAgentResponse
	GetBody() *CloudBatchUpdateAgentResponseBody
}

type CloudBatchUpdateAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudBatchUpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudBatchUpdateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchUpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudBatchUpdateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudBatchUpdateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudBatchUpdateAgentResponse) GetBody() *CloudBatchUpdateAgentResponseBody {
	return s.Body
}

func (s *CloudBatchUpdateAgentResponse) SetHeaders(v map[string]*string) *CloudBatchUpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudBatchUpdateAgentResponse) SetStatusCode(v int32) *CloudBatchUpdateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudBatchUpdateAgentResponse) SetBody(v *CloudBatchUpdateAgentResponseBody) *CloudBatchUpdateAgentResponse {
	s.Body = v
	return s
}

func (s *CloudBatchUpdateAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
