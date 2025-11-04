// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *DropSearchLibResponseBody) *DropSearchLibResponse
	GetBody() *DropSearchLibResponseBody
}

type DropSearchLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DropSearchLibResponse) GoString() string {
	return s.String()
}

func (s *DropSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropSearchLibResponse) GetBody() *DropSearchLibResponseBody {
	return s.Body
}

func (s *DropSearchLibResponse) SetHeaders(v map[string]*string) *DropSearchLibResponse {
	s.Headers = v
	return s
}

func (s *DropSearchLibResponse) SetStatusCode(v int32) *DropSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSearchLibResponse) SetBody(v *DropSearchLibResponseBody) *DropSearchLibResponse {
	s.Body = v
	return s
}

func (s *DropSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
