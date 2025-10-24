// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDefenseResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDefenseResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDefenseResourceResponseBody) *DeleteDefenseResourceResponse
	GetBody() *DeleteDefenseResourceResponseBody
}

type DeleteDefenseResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDefenseResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDefenseResourceResponse) GetBody() *DeleteDefenseResourceResponseBody {
	return s.Body
}

func (s *DeleteDefenseResourceResponse) SetHeaders(v map[string]*string) *DeleteDefenseResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseResourceResponse) SetStatusCode(v int32) *DeleteDefenseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseResourceResponse) SetBody(v *DeleteDefenseResourceResponseBody) *DeleteDefenseResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteDefenseResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
