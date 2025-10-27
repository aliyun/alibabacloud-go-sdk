// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkFlowResponseBody) *GetWorkFlowResponse
	GetBody() *GetWorkFlowResponseBody
}

type GetWorkFlowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponse) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkFlowResponse) GetBody() *GetWorkFlowResponseBody {
	return s.Body
}

func (s *GetWorkFlowResponse) SetHeaders(v map[string]*string) *GetWorkFlowResponse {
	s.Headers = v
	return s
}

func (s *GetWorkFlowResponse) SetStatusCode(v int32) *GetWorkFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkFlowResponse) SetBody(v *GetWorkFlowResponseBody) *GetWorkFlowResponse {
	s.Body = v
	return s
}

func (s *GetWorkFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
