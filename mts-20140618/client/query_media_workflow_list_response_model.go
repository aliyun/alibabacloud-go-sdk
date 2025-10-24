// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaWorkflowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaWorkflowListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaWorkflowListResponseBody) *QueryMediaWorkflowListResponse
	GetBody() *QueryMediaWorkflowListResponseBody
}

type QueryMediaWorkflowListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaWorkflowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaWorkflowListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaWorkflowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaWorkflowListResponse) GetBody() *QueryMediaWorkflowListResponseBody {
	return s.Body
}

func (s *QueryMediaWorkflowListResponse) SetHeaders(v map[string]*string) *QueryMediaWorkflowListResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaWorkflowListResponse) SetStatusCode(v int32) *QueryMediaWorkflowListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaWorkflowListResponse) SetBody(v *QueryMediaWorkflowListResponseBody) *QueryMediaWorkflowListResponse {
	s.Body = v
	return s
}

func (s *QueryMediaWorkflowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
