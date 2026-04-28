// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFacegroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFacegroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFacegroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFacegroupResponseBody) *UpdateFacegroupResponse
	GetBody() *UpdateFacegroupResponseBody
}

type UpdateFacegroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFacegroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFacegroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFacegroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFacegroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFacegroupResponse) GetBody() *UpdateFacegroupResponseBody {
	return s.Body
}

func (s *UpdateFacegroupResponse) SetHeaders(v map[string]*string) *UpdateFacegroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFacegroupResponse) SetStatusCode(v int32) *UpdateFacegroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFacegroupResponse) SetBody(v *UpdateFacegroupResponseBody) *UpdateFacegroupResponse {
	s.Body = v
	return s
}

func (s *UpdateFacegroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
