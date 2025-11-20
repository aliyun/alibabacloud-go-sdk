// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNasFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachNasFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachNasFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *DetachNasFileSystemResponseBody) *DetachNasFileSystemResponse
	GetBody() *DetachNasFileSystemResponseBody
}

type DetachNasFileSystemResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachNasFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachNasFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachNasFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachNasFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachNasFileSystemResponse) GetBody() *DetachNasFileSystemResponseBody {
	return s.Body
}

func (s *DetachNasFileSystemResponse) SetHeaders(v map[string]*string) *DetachNasFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DetachNasFileSystemResponse) SetStatusCode(v int32) *DetachNasFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachNasFileSystemResponse) SetBody(v *DetachNasFileSystemResponseBody) *DetachNasFileSystemResponse {
	s.Body = v
	return s
}

func (s *DetachNasFileSystemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
