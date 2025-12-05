// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultKmsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultKmsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultKmsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetDefaultKmsInstanceResponseBody) *GetDefaultKmsInstanceResponse
	GetBody() *GetDefaultKmsInstanceResponseBody
}

type GetDefaultKmsInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultKmsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultKmsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultKmsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultKmsInstanceResponse) GetBody() *GetDefaultKmsInstanceResponseBody {
	return s.Body
}

func (s *GetDefaultKmsInstanceResponse) SetHeaders(v map[string]*string) *GetDefaultKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultKmsInstanceResponse) SetStatusCode(v int32) *GetDefaultKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultKmsInstanceResponse) SetBody(v *GetDefaultKmsInstanceResponseBody) *GetDefaultKmsInstanceResponse {
	s.Body = v
	return s
}

func (s *GetDefaultKmsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
