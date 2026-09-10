// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaOssTempKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaOssTempKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaOssTempKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaOssTempKeyResponseBody) *GetMetaOssTempKeyResponse
	GetBody() *GetMetaOssTempKeyResponseBody
}

type GetMetaOssTempKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaOssTempKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaOssTempKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaOssTempKeyResponse) GoString() string {
	return s.String()
}

func (s *GetMetaOssTempKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaOssTempKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaOssTempKeyResponse) GetBody() *GetMetaOssTempKeyResponseBody {
	return s.Body
}

func (s *GetMetaOssTempKeyResponse) SetHeaders(v map[string]*string) *GetMetaOssTempKeyResponse {
	s.Headers = v
	return s
}

func (s *GetMetaOssTempKeyResponse) SetStatusCode(v int32) *GetMetaOssTempKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaOssTempKeyResponse) SetBody(v *GetMetaOssTempKeyResponseBody) *GetMetaOssTempKeyResponse {
	s.Body = v
	return s
}

func (s *GetMetaOssTempKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
