// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceSharePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateResourceSharePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateResourceSharePermissionResponse
	GetStatusCode() *int32
	SetBody(v *AssociateResourceSharePermissionResponseBody) *AssociateResourceSharePermissionResponse
	GetBody() *AssociateResourceSharePermissionResponseBody
}

type AssociateResourceSharePermissionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResourceSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateResourceSharePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateResourceSharePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateResourceSharePermissionResponse) GetBody() *AssociateResourceSharePermissionResponseBody {
	return s.Body
}

func (s *AssociateResourceSharePermissionResponse) SetHeaders(v map[string]*string) *AssociateResourceSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *AssociateResourceSharePermissionResponse) SetStatusCode(v int32) *AssociateResourceSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResourceSharePermissionResponse) SetBody(v *AssociateResourceSharePermissionResponseBody) *AssociateResourceSharePermissionResponse {
	s.Body = v
	return s
}

func (s *AssociateResourceSharePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
