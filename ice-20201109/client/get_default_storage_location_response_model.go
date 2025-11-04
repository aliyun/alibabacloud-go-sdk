// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultStorageLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultStorageLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultStorageLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetDefaultStorageLocationResponseBody) *GetDefaultStorageLocationResponse
	GetBody() *GetDefaultStorageLocationResponseBody
}

type GetDefaultStorageLocationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultStorageLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultStorageLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultStorageLocationResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultStorageLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultStorageLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultStorageLocationResponse) GetBody() *GetDefaultStorageLocationResponseBody {
	return s.Body
}

func (s *GetDefaultStorageLocationResponse) SetHeaders(v map[string]*string) *GetDefaultStorageLocationResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultStorageLocationResponse) SetStatusCode(v int32) *GetDefaultStorageLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultStorageLocationResponse) SetBody(v *GetDefaultStorageLocationResponseBody) *GetDefaultStorageLocationResponse {
	s.Body = v
	return s
}

func (s *GetDefaultStorageLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
