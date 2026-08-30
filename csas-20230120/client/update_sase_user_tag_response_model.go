// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSaseUserTagResponseBody) *UpdateSaseUserTagResponse
	GetBody() *UpdateSaseUserTagResponseBody
}

type UpdateSaseUserTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSaseUserTagResponse) GetBody() *UpdateSaseUserTagResponseBody {
	return s.Body
}

func (s *UpdateSaseUserTagResponse) SetHeaders(v map[string]*string) *UpdateSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateSaseUserTagResponse) SetStatusCode(v int32) *UpdateSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSaseUserTagResponse) SetBody(v *UpdateSaseUserTagResponseBody) *UpdateSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *UpdateSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
