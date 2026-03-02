// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcVersionNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcVersionNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcVersionNumbersResponse
	GetStatusCode() *int32
	SetBody(v *PbcVersionListResult) *ListPbcVersionNumbersResponse
	GetBody() *PbcVersionListResult
}

type ListPbcVersionNumbersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcVersionListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcVersionNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcVersionNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListPbcVersionNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcVersionNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcVersionNumbersResponse) GetBody() *PbcVersionListResult {
	return s.Body
}

func (s *ListPbcVersionNumbersResponse) SetHeaders(v map[string]*string) *ListPbcVersionNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListPbcVersionNumbersResponse) SetStatusCode(v int32) *ListPbcVersionNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcVersionNumbersResponse) SetBody(v *PbcVersionListResult) *ListPbcVersionNumbersResponse {
	s.Body = v
	return s
}

func (s *ListPbcVersionNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
