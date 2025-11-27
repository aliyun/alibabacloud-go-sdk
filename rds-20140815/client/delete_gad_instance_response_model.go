// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGadInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGadInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGadInstanceResponseBody) *DeleteGadInstanceResponse
	GetBody() *DeleteGadInstanceResponseBody
}

type DeleteGadInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGadInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGadInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGadInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGadInstanceResponse) GetBody() *DeleteGadInstanceResponseBody {
	return s.Body
}

func (s *DeleteGadInstanceResponse) SetHeaders(v map[string]*string) *DeleteGadInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGadInstanceResponse) SetStatusCode(v int32) *DeleteGadInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGadInstanceResponse) SetBody(v *DeleteGadInstanceResponseBody) *DeleteGadInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteGadInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
