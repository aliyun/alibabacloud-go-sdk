// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrashFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrashFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrashFileResponse
	GetStatusCode() *int32
	SetBody(v *TrashFileResponseBody) *TrashFileResponse
	GetBody() *TrashFileResponseBody
}

type TrashFileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrashFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrashFileResponse) String() string {
	return dara.Prettify(s)
}

func (s TrashFileResponse) GoString() string {
	return s.String()
}

func (s *TrashFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrashFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrashFileResponse) GetBody() *TrashFileResponseBody {
	return s.Body
}

func (s *TrashFileResponse) SetHeaders(v map[string]*string) *TrashFileResponse {
	s.Headers = v
	return s
}

func (s *TrashFileResponse) SetStatusCode(v int32) *TrashFileResponse {
	s.StatusCode = &v
	return s
}

func (s *TrashFileResponse) SetBody(v *TrashFileResponseBody) *TrashFileResponse {
	s.Body = v
	return s
}

func (s *TrashFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
