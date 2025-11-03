// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowLogResponseBody) *CreateFlowLogResponse
	GetBody() *CreateFlowLogResponseBody
}

type CreateFlowLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowLogResponse) GetBody() *CreateFlowLogResponseBody {
	return s.Body
}

func (s *CreateFlowLogResponse) SetHeaders(v map[string]*string) *CreateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowLogResponse) SetStatusCode(v int32) *CreateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowLogResponse) SetBody(v *CreateFlowLogResponseBody) *CreateFlowLogResponse {
	s.Body = v
	return s
}

func (s *CreateFlowLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
