// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProhibitedTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProhibitedTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProhibitedTagResponseBody) *UpdateProhibitedTagResponse
	GetBody() *UpdateProhibitedTagResponseBody
}

type UpdateProhibitedTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProhibitedTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProhibitedTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProhibitedTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProhibitedTagResponse) GetBody() *UpdateProhibitedTagResponseBody {
	return s.Body
}

func (s *UpdateProhibitedTagResponse) SetHeaders(v map[string]*string) *UpdateProhibitedTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateProhibitedTagResponse) SetStatusCode(v int32) *UpdateProhibitedTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProhibitedTagResponse) SetBody(v *UpdateProhibitedTagResponseBody) *UpdateProhibitedTagResponse {
	s.Body = v
	return s
}

func (s *UpdateProhibitedTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
