// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateACLResponse
	GetStatusCode() *int32
	SetBody(v *AssociateACLResponseBody) *AssociateACLResponse
	GetBody() *AssociateACLResponseBody
}

type AssociateACLResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateACLResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateACLResponse) GoString() string {
	return s.String()
}

func (s *AssociateACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateACLResponse) GetBody() *AssociateACLResponseBody {
	return s.Body
}

func (s *AssociateACLResponse) SetHeaders(v map[string]*string) *AssociateACLResponse {
	s.Headers = v
	return s
}

func (s *AssociateACLResponse) SetStatusCode(v int32) *AssociateACLResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateACLResponse) SetBody(v *AssociateACLResponseBody) *AssociateACLResponse {
	s.Body = v
	return s
}

func (s *AssociateACLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
