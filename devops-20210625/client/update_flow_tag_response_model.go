// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowTagResponseBody) *UpdateFlowTagResponse
	GetBody() *UpdateFlowTagResponseBody
}

type UpdateFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowTagResponse) GetBody() *UpdateFlowTagResponseBody {
	return s.Body
}

func (s *UpdateFlowTagResponse) SetHeaders(v map[string]*string) *UpdateFlowTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowTagResponse) SetStatusCode(v int32) *UpdateFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowTagResponse) SetBody(v *UpdateFlowTagResponseBody) *UpdateFlowTagResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
