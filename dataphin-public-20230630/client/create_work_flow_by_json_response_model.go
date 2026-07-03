// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkFlowByJsonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkFlowByJsonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkFlowByJsonResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkFlowByJsonResponseBody) *CreateWorkFlowByJsonResponse
	GetBody() *CreateWorkFlowByJsonResponseBody
}

type CreateWorkFlowByJsonResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkFlowByJsonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkFlowByJsonResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkFlowByJsonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkFlowByJsonResponse) GetBody() *CreateWorkFlowByJsonResponseBody {
	return s.Body
}

func (s *CreateWorkFlowByJsonResponse) SetHeaders(v map[string]*string) *CreateWorkFlowByJsonResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkFlowByJsonResponse) SetStatusCode(v int32) *CreateWorkFlowByJsonResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkFlowByJsonResponse) SetBody(v *CreateWorkFlowByJsonResponseBody) *CreateWorkFlowByJsonResponse {
	s.Body = v
	return s
}

func (s *CreateWorkFlowByJsonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
