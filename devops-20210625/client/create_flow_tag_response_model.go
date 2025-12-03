// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowTagResponseBody) *CreateFlowTagResponse
	GetBody() *CreateFlowTagResponseBody
}

type CreateFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowTagResponse) GetBody() *CreateFlowTagResponseBody {
	return s.Body
}

func (s *CreateFlowTagResponse) SetHeaders(v map[string]*string) *CreateFlowTagResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowTagResponse) SetStatusCode(v int32) *CreateFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowTagResponse) SetBody(v *CreateFlowTagResponseBody) *CreateFlowTagResponse {
	s.Body = v
	return s
}

func (s *CreateFlowTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
