// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiServerEipResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiServerEipResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiServerEipResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiServerEipResourceResponseBody) *ModifyApiServerEipResourceResponse
	GetBody() *ModifyApiServerEipResourceResponseBody
}

type ModifyApiServerEipResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiServerEipResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiServerEipResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiServerEipResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiServerEipResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiServerEipResourceResponse) GetBody() *ModifyApiServerEipResourceResponseBody {
	return s.Body
}

func (s *ModifyApiServerEipResourceResponse) SetHeaders(v map[string]*string) *ModifyApiServerEipResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiServerEipResourceResponse) SetStatusCode(v int32) *ModifyApiServerEipResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiServerEipResourceResponse) SetBody(v *ModifyApiServerEipResourceResponseBody) *ModifyApiServerEipResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyApiServerEipResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
