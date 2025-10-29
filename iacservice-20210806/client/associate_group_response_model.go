// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateGroupResponse
	GetStatusCode() *int32
	SetBody(v *AssociateGroupResponseBody) *AssociateGroupResponse
	GetBody() *AssociateGroupResponseBody
}

type AssociateGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateGroupResponse) GoString() string {
	return s.String()
}

func (s *AssociateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateGroupResponse) GetBody() *AssociateGroupResponseBody {
	return s.Body
}

func (s *AssociateGroupResponse) SetHeaders(v map[string]*string) *AssociateGroupResponse {
	s.Headers = v
	return s
}

func (s *AssociateGroupResponse) SetStatusCode(v int32) *AssociateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateGroupResponse) SetBody(v *AssociateGroupResponseBody) *AssociateGroupResponse {
	s.Body = v
	return s
}

func (s *AssociateGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
