// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentCnoAndNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryAgentCnoAndNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryAgentCnoAndNameResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryAgentCnoAndNameResponseBody) *CloudQueryAgentCnoAndNameResponse
	GetBody() *CloudQueryAgentCnoAndNameResponseBody
}

type CloudQueryAgentCnoAndNameResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryAgentCnoAndNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryAgentCnoAndNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentCnoAndNameResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentCnoAndNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryAgentCnoAndNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryAgentCnoAndNameResponse) GetBody() *CloudQueryAgentCnoAndNameResponseBody {
	return s.Body
}

func (s *CloudQueryAgentCnoAndNameResponse) SetHeaders(v map[string]*string) *CloudQueryAgentCnoAndNameResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponse) SetStatusCode(v int32) *CloudQueryAgentCnoAndNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponse) SetBody(v *CloudQueryAgentCnoAndNameResponseBody) *CloudQueryAgentCnoAndNameResponse {
	s.Body = v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
