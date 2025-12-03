// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemWorkFlowInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemWorkFlowInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemWorkFlowInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemWorkFlowInfoResponseBody) *GetWorkItemWorkFlowInfoResponse
	GetBody() *GetWorkItemWorkFlowInfoResponseBody
}

type GetWorkItemWorkFlowInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemWorkFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemWorkFlowInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemWorkFlowInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemWorkFlowInfoResponse) GetBody() *GetWorkItemWorkFlowInfoResponseBody {
	return s.Body
}

func (s *GetWorkItemWorkFlowInfoResponse) SetHeaders(v map[string]*string) *GetWorkItemWorkFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponse) SetStatusCode(v int32) *GetWorkItemWorkFlowInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponse) SetBody(v *GetWorkItemWorkFlowInfoResponseBody) *GetWorkItemWorkFlowInfoResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
