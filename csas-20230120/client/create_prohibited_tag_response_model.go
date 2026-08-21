// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProhibitedTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProhibitedTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateProhibitedTagResponseBody) *CreateProhibitedTagResponse
	GetBody() *CreateProhibitedTagResponseBody
}

type CreateProhibitedTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProhibitedTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProhibitedTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedTagResponse) GoString() string {
	return s.String()
}

func (s *CreateProhibitedTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProhibitedTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProhibitedTagResponse) GetBody() *CreateProhibitedTagResponseBody {
	return s.Body
}

func (s *CreateProhibitedTagResponse) SetHeaders(v map[string]*string) *CreateProhibitedTagResponse {
	s.Headers = v
	return s
}

func (s *CreateProhibitedTagResponse) SetStatusCode(v int32) *CreateProhibitedTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProhibitedTagResponse) SetBody(v *CreateProhibitedTagResponseBody) *CreateProhibitedTagResponse {
	s.Body = v
	return s
}

func (s *CreateProhibitedTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
