// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQosAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQosAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQosAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetQosAttributeResponseBody) *GetQosAttributeResponse
	GetBody() *GetQosAttributeResponseBody
}

type GetQosAttributeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQosAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQosAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQosAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetQosAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQosAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQosAttributeResponse) GetBody() *GetQosAttributeResponseBody {
	return s.Body
}

func (s *GetQosAttributeResponse) SetHeaders(v map[string]*string) *GetQosAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetQosAttributeResponse) SetStatusCode(v int32) *GetQosAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQosAttributeResponse) SetBody(v *GetQosAttributeResponseBody) *GetQosAttributeResponse {
	s.Body = v
	return s
}

func (s *GetQosAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
