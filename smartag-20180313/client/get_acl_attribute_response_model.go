// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAclAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAclAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetAclAttributeResponseBody) *GetAclAttributeResponse
	GetBody() *GetAclAttributeResponseBody
}

type GetAclAttributeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAclAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetAclAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAclAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAclAttributeResponse) GetBody() *GetAclAttributeResponseBody {
	return s.Body
}

func (s *GetAclAttributeResponse) SetHeaders(v map[string]*string) *GetAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetAclAttributeResponse) SetStatusCode(v int32) *GetAclAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAclAttributeResponse) SetBody(v *GetAclAttributeResponseBody) *GetAclAttributeResponse {
	s.Body = v
	return s
}

func (s *GetAclAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
