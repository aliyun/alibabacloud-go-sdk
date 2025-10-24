// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaWorkflowResponseBody) *SearchMediaWorkflowResponse
	GetBody() *SearchMediaWorkflowResponseBody
}

type SearchMediaWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaWorkflowResponse) GetBody() *SearchMediaWorkflowResponseBody {
	return s.Body
}

func (s *SearchMediaWorkflowResponse) SetHeaders(v map[string]*string) *SearchMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaWorkflowResponse) SetStatusCode(v int32) *SearchMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaWorkflowResponse) SetBody(v *SearchMediaWorkflowResponseBody) *SearchMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *SearchMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
