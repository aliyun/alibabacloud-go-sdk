// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueCreateStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinueCreateStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinueCreateStackResponse
	GetStatusCode() *int32
	SetBody(v *ContinueCreateStackResponseBody) *ContinueCreateStackResponse
	GetBody() *ContinueCreateStackResponseBody
}

type ContinueCreateStackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueCreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueCreateStackResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinueCreateStackResponse) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinueCreateStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinueCreateStackResponse) GetBody() *ContinueCreateStackResponseBody {
	return s.Body
}

func (s *ContinueCreateStackResponse) SetHeaders(v map[string]*string) *ContinueCreateStackResponse {
	s.Headers = v
	return s
}

func (s *ContinueCreateStackResponse) SetStatusCode(v int32) *ContinueCreateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueCreateStackResponse) SetBody(v *ContinueCreateStackResponseBody) *ContinueCreateStackResponse {
	s.Body = v
	return s
}

func (s *ContinueCreateStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
