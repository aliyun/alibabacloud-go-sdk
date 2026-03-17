// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateQosResponse
	GetStatusCode() *int32
	SetBody(v *AssociateQosResponseBody) *AssociateQosResponse
	GetBody() *AssociateQosResponseBody
}

type AssociateQosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateQosResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateQosResponse) GoString() string {
	return s.String()
}

func (s *AssociateQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateQosResponse) GetBody() *AssociateQosResponseBody {
	return s.Body
}

func (s *AssociateQosResponse) SetHeaders(v map[string]*string) *AssociateQosResponse {
	s.Headers = v
	return s
}

func (s *AssociateQosResponse) SetStatusCode(v int32) *AssociateQosResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateQosResponse) SetBody(v *AssociateQosResponseBody) *AssociateQosResponse {
	s.Body = v
	return s
}

func (s *AssociateQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
