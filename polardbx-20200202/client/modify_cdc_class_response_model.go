// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdcClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdcClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdcClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdcClassResponseBody) *ModifyCdcClassResponse
	GetBody() *ModifyCdcClassResponseBody
}

type ModifyCdcClassResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdcClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdcClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdcClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdcClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdcClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdcClassResponse) GetBody() *ModifyCdcClassResponseBody {
	return s.Body
}

func (s *ModifyCdcClassResponse) SetHeaders(v map[string]*string) *ModifyCdcClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdcClassResponse) SetStatusCode(v int32) *ModifyCdcClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdcClassResponse) SetBody(v *ModifyCdcClassResponseBody) *ModifyCdcClassResponse {
	s.Body = v
	return s
}

func (s *ModifyCdcClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
