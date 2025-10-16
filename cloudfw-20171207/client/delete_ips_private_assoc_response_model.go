// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsPrivateAssocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpsPrivateAssocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpsPrivateAssocResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpsPrivateAssocResponseBody) *DeleteIpsPrivateAssocResponse
	GetBody() *DeleteIpsPrivateAssocResponseBody
}

type DeleteIpsPrivateAssocResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpsPrivateAssocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpsPrivateAssocResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsPrivateAssocResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpsPrivateAssocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpsPrivateAssocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpsPrivateAssocResponse) GetBody() *DeleteIpsPrivateAssocResponseBody {
	return s.Body
}

func (s *DeleteIpsPrivateAssocResponse) SetHeaders(v map[string]*string) *DeleteIpsPrivateAssocResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpsPrivateAssocResponse) SetStatusCode(v int32) *DeleteIpsPrivateAssocResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpsPrivateAssocResponse) SetBody(v *DeleteIpsPrivateAssocResponseBody) *DeleteIpsPrivateAssocResponse {
	s.Body = v
	return s
}

func (s *DeleteIpsPrivateAssocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
