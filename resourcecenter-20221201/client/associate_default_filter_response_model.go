// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDefaultFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateDefaultFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateDefaultFilterResponse
	GetStatusCode() *int32
	SetBody(v *AssociateDefaultFilterResponseBody) *AssociateDefaultFilterResponse
	GetBody() *AssociateDefaultFilterResponseBody
}

type AssociateDefaultFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateDefaultFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateDefaultFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateDefaultFilterResponse) GetBody() *AssociateDefaultFilterResponseBody {
	return s.Body
}

func (s *AssociateDefaultFilterResponse) SetHeaders(v map[string]*string) *AssociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *AssociateDefaultFilterResponse) SetStatusCode(v int32) *AssociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateDefaultFilterResponse) SetBody(v *AssociateDefaultFilterResponseBody) *AssociateDefaultFilterResponse {
	s.Body = v
	return s
}

func (s *AssociateDefaultFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
