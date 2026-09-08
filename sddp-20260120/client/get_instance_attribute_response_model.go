// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAttributeResponseBody) *GetInstanceAttributeResponse
	GetBody() *GetInstanceAttributeResponseBody
}

type GetInstanceAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAttributeResponse) GetBody() *GetInstanceAttributeResponseBody {
	return s.Body
}

func (s *GetInstanceAttributeResponse) SetHeaders(v map[string]*string) *GetInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAttributeResponse) SetStatusCode(v int32) *GetInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAttributeResponse) SetBody(v *GetInstanceAttributeResponseBody) *GetInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
