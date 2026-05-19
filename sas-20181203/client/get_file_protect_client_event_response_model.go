// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectClientEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectClientEventResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectClientEventResponseBody) *GetFileProtectClientEventResponse
	GetBody() *GetFileProtectClientEventResponseBody
}

type GetFileProtectClientEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectClientEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectClientEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectClientEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectClientEventResponse) GetBody() *GetFileProtectClientEventResponseBody {
	return s.Body
}

func (s *GetFileProtectClientEventResponse) SetHeaders(v map[string]*string) *GetFileProtectClientEventResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectClientEventResponse) SetStatusCode(v int32) *GetFileProtectClientEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectClientEventResponse) SetBody(v *GetFileProtectClientEventResponseBody) *GetFileProtectClientEventResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectClientEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
