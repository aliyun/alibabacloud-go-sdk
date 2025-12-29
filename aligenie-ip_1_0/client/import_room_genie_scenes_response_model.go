// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomGenieScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportRoomGenieScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportRoomGenieScenesResponse
	GetStatusCode() *int32
	SetBody(v *ImportRoomGenieScenesResponseBody) *ImportRoomGenieScenesResponse
	GetBody() *ImportRoomGenieScenesResponseBody
}

type ImportRoomGenieScenesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportRoomGenieScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportRoomGenieScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesResponse) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportRoomGenieScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportRoomGenieScenesResponse) GetBody() *ImportRoomGenieScenesResponseBody {
	return s.Body
}

func (s *ImportRoomGenieScenesResponse) SetHeaders(v map[string]*string) *ImportRoomGenieScenesResponse {
	s.Headers = v
	return s
}

func (s *ImportRoomGenieScenesResponse) SetStatusCode(v int32) *ImportRoomGenieScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomGenieScenesResponse) SetBody(v *ImportRoomGenieScenesResponseBody) *ImportRoomGenieScenesResponse {
	s.Body = v
	return s
}

func (s *ImportRoomGenieScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
