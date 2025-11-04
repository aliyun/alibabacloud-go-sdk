// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultStorageLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultStorageLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultStorageLocationResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultStorageLocationResponseBody) *SetDefaultStorageLocationResponse
	GetBody() *SetDefaultStorageLocationResponseBody
}

type SetDefaultStorageLocationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultStorageLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultStorageLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultStorageLocationResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultStorageLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultStorageLocationResponse) GetBody() *SetDefaultStorageLocationResponseBody {
	return s.Body
}

func (s *SetDefaultStorageLocationResponse) SetHeaders(v map[string]*string) *SetDefaultStorageLocationResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultStorageLocationResponse) SetStatusCode(v int32) *SetDefaultStorageLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultStorageLocationResponse) SetBody(v *SetDefaultStorageLocationResponseBody) *SetDefaultStorageLocationResponse {
	s.Body = v
	return s
}

func (s *SetDefaultStorageLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
