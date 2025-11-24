// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVmMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVmMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetVmMetaResponseBody) *GetVmMetaResponse
	GetBody() *GetVmMetaResponseBody
}

type GetVmMetaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVmMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVmMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVmMetaResponse) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVmMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVmMetaResponse) GetBody() *GetVmMetaResponseBody {
	return s.Body
}

func (s *GetVmMetaResponse) SetHeaders(v map[string]*string) *GetVmMetaResponse {
	s.Headers = v
	return s
}

func (s *GetVmMetaResponse) SetStatusCode(v int32) *GetVmMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVmMetaResponse) SetBody(v *GetVmMetaResponseBody) *GetVmMetaResponse {
	s.Body = v
	return s
}

func (s *GetVmMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
