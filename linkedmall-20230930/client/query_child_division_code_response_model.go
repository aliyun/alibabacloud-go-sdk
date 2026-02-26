// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChildDivisionCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryChildDivisionCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryChildDivisionCodeResponse
	GetStatusCode() *int32
	SetBody(v *DivisionPageResult) *QueryChildDivisionCodeResponse
	GetBody() *DivisionPageResult
}

type QueryChildDivisionCodeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DivisionPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChildDivisionCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryChildDivisionCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryChildDivisionCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryChildDivisionCodeResponse) GetBody() *DivisionPageResult {
	return s.Body
}

func (s *QueryChildDivisionCodeResponse) SetHeaders(v map[string]*string) *QueryChildDivisionCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryChildDivisionCodeResponse) SetStatusCode(v int32) *QueryChildDivisionCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChildDivisionCodeResponse) SetBody(v *DivisionPageResult) *QueryChildDivisionCodeResponse {
	s.Body = v
	return s
}

func (s *QueryChildDivisionCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
