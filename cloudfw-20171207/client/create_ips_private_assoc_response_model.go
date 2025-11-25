// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsPrivateAssocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpsPrivateAssocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpsPrivateAssocResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpsPrivateAssocResponseBody) *CreateIpsPrivateAssocResponse
	GetBody() *CreateIpsPrivateAssocResponseBody
}

type CreateIpsPrivateAssocResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpsPrivateAssocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpsPrivateAssocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsPrivateAssocResponse) GoString() string {
	return s.String()
}

func (s *CreateIpsPrivateAssocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpsPrivateAssocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpsPrivateAssocResponse) GetBody() *CreateIpsPrivateAssocResponseBody {
	return s.Body
}

func (s *CreateIpsPrivateAssocResponse) SetHeaders(v map[string]*string) *CreateIpsPrivateAssocResponse {
	s.Headers = v
	return s
}

func (s *CreateIpsPrivateAssocResponse) SetStatusCode(v int32) *CreateIpsPrivateAssocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpsPrivateAssocResponse) SetBody(v *CreateIpsPrivateAssocResponseBody) *CreateIpsPrivateAssocResponse {
	s.Body = v
	return s
}

func (s *CreateIpsPrivateAssocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
