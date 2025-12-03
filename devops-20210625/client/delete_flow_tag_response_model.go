// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowTagResponseBody) *DeleteFlowTagResponse
	GetBody() *DeleteFlowTagResponseBody
}

type DeleteFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowTagResponse) GetBody() *DeleteFlowTagResponseBody {
	return s.Body
}

func (s *DeleteFlowTagResponse) SetHeaders(v map[string]*string) *DeleteFlowTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowTagResponse) SetStatusCode(v int32) *DeleteFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowTagResponse) SetBody(v *DeleteFlowTagResponseBody) *DeleteFlowTagResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
