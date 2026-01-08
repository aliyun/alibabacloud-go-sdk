// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowVersionResponseBody) *CreateFlowVersionResponse
	GetBody() *CreateFlowVersionResponseBody
}

type CreateFlowVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowVersionResponse) GetBody() *CreateFlowVersionResponseBody {
	return s.Body
}

func (s *CreateFlowVersionResponse) SetHeaders(v map[string]*string) *CreateFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowVersionResponse) SetStatusCode(v int32) *CreateFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowVersionResponse) SetBody(v *CreateFlowVersionResponseBody) *CreateFlowVersionResponse {
	s.Body = v
	return s
}

func (s *CreateFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
