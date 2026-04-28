// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDriveResponse
	GetStatusCode() *int32
	SetBody(v *SearchDriveResponseBody) *SearchDriveResponse
	GetBody() *SearchDriveResponseBody
}

type SearchDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDriveResponse) GoString() string {
	return s.String()
}

func (s *SearchDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDriveResponse) GetBody() *SearchDriveResponseBody {
	return s.Body
}

func (s *SearchDriveResponse) SetHeaders(v map[string]*string) *SearchDriveResponse {
	s.Headers = v
	return s
}

func (s *SearchDriveResponse) SetStatusCode(v int32) *SearchDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDriveResponse) SetBody(v *SearchDriveResponseBody) *SearchDriveResponse {
	s.Body = v
	return s
}

func (s *SearchDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
