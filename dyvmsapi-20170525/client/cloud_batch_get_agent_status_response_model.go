// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchGetAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudBatchGetAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudBatchGetAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *CloudBatchGetAgentStatusResponseBody) *CloudBatchGetAgentStatusResponse
	GetBody() *CloudBatchGetAgentStatusResponseBody
}

type CloudBatchGetAgentStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudBatchGetAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudBatchGetAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchGetAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *CloudBatchGetAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudBatchGetAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudBatchGetAgentStatusResponse) GetBody() *CloudBatchGetAgentStatusResponseBody {
	return s.Body
}

func (s *CloudBatchGetAgentStatusResponse) SetHeaders(v map[string]*string) *CloudBatchGetAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *CloudBatchGetAgentStatusResponse) SetStatusCode(v int32) *CloudBatchGetAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudBatchGetAgentStatusResponse) SetBody(v *CloudBatchGetAgentStatusResponseBody) *CloudBatchGetAgentStatusResponse {
	s.Body = v
	return s
}

func (s *CloudBatchGetAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
