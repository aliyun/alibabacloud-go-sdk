// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameSourceResponse
	GetStatusCode() *int32
	SetBody(v *RenameSourceResponseBody) *RenameSourceResponse
	GetBody() *RenameSourceResponseBody
}

type RenameSourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameSourceResponse) GoString() string {
	return s.String()
}

func (s *RenameSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameSourceResponse) GetBody() *RenameSourceResponseBody {
	return s.Body
}

func (s *RenameSourceResponse) SetHeaders(v map[string]*string) *RenameSourceResponse {
	s.Headers = v
	return s
}

func (s *RenameSourceResponse) SetStatusCode(v int32) *RenameSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameSourceResponse) SetBody(v *RenameSourceResponseBody) *RenameSourceResponse {
	s.Body = v
	return s
}

func (s *RenameSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
