// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateParameterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateParameterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateParameterSetResponse
	GetStatusCode() *int32
	SetBody(v *AssociateParameterSetResponseBody) *AssociateParameterSetResponse
	GetBody() *AssociateParameterSetResponseBody
}

type AssociateParameterSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateParameterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateParameterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateParameterSetResponse) GetBody() *AssociateParameterSetResponseBody {
	return s.Body
}

func (s *AssociateParameterSetResponse) SetHeaders(v map[string]*string) *AssociateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *AssociateParameterSetResponse) SetStatusCode(v int32) *AssociateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateParameterSetResponse) SetBody(v *AssociateParameterSetResponseBody) *AssociateParameterSetResponse {
	s.Body = v
	return s
}

func (s *AssociateParameterSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
