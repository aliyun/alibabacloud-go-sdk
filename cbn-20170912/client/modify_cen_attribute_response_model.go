// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCenAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCenAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCenAttributeResponseBody) *ModifyCenAttributeResponse
	GetBody() *ModifyCenAttributeResponseBody
}

type ModifyCenAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCenAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCenAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCenAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCenAttributeResponse) GetBody() *ModifyCenAttributeResponseBody {
	return s.Body
}

func (s *ModifyCenAttributeResponse) SetHeaders(v map[string]*string) *ModifyCenAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenAttributeResponse) SetStatusCode(v int32) *ModifyCenAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCenAttributeResponse) SetBody(v *ModifyCenAttributeResponseBody) *ModifyCenAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyCenAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
