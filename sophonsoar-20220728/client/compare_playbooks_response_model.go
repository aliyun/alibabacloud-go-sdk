// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComparePlaybooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ComparePlaybooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ComparePlaybooksResponse
	GetStatusCode() *int32
	SetBody(v *ComparePlaybooksResponseBody) *ComparePlaybooksResponse
	GetBody() *ComparePlaybooksResponseBody
}

type ComparePlaybooksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComparePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ComparePlaybooksResponse) String() string {
	return dara.Prettify(s)
}

func (s ComparePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ComparePlaybooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ComparePlaybooksResponse) GetBody() *ComparePlaybooksResponseBody {
	return s.Body
}

func (s *ComparePlaybooksResponse) SetHeaders(v map[string]*string) *ComparePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *ComparePlaybooksResponse) SetStatusCode(v int32) *ComparePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ComparePlaybooksResponse) SetBody(v *ComparePlaybooksResponseBody) *ComparePlaybooksResponse {
	s.Body = v
	return s
}

func (s *ComparePlaybooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
