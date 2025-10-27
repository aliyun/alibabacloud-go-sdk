// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectEventResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectEventResponseBody) *GetFileProtectEventResponse
	GetBody() *GetFileProtectEventResponseBody
}

type GetFileProtectEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectEventResponse) GetBody() *GetFileProtectEventResponseBody {
	return s.Body
}

func (s *GetFileProtectEventResponse) SetHeaders(v map[string]*string) *GetFileProtectEventResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectEventResponse) SetStatusCode(v int32) *GetFileProtectEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectEventResponse) SetBody(v *GetFileProtectEventResponseBody) *GetFileProtectEventResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
