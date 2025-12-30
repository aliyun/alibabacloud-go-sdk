// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *AlterSearchLibResponseBody) *AlterSearchLibResponse
	GetBody() *AlterSearchLibResponseBody
}

type AlterSearchLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlterSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlterSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchLibResponse) GoString() string {
	return s.String()
}

func (s *AlterSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterSearchLibResponse) GetBody() *AlterSearchLibResponseBody {
	return s.Body
}

func (s *AlterSearchLibResponse) SetHeaders(v map[string]*string) *AlterSearchLibResponse {
	s.Headers = v
	return s
}

func (s *AlterSearchLibResponse) SetStatusCode(v int32) *AlterSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterSearchLibResponse) SetBody(v *AlterSearchLibResponseBody) *AlterSearchLibResponse {
	s.Body = v
	return s
}

func (s *AlterSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
