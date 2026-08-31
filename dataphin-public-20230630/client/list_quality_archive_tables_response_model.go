// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityArchiveTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityArchiveTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityArchiveTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityArchiveTablesResponseBody) *ListQualityArchiveTablesResponse
	GetBody() *ListQualityArchiveTablesResponseBody
}

type ListQualityArchiveTablesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityArchiveTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityArchiveTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesResponse) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityArchiveTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityArchiveTablesResponse) GetBody() *ListQualityArchiveTablesResponseBody {
	return s.Body
}

func (s *ListQualityArchiveTablesResponse) SetHeaders(v map[string]*string) *ListQualityArchiveTablesResponse {
	s.Headers = v
	return s
}

func (s *ListQualityArchiveTablesResponse) SetStatusCode(v int32) *ListQualityArchiveTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityArchiveTablesResponse) SetBody(v *ListQualityArchiveTablesResponseBody) *ListQualityArchiveTablesResponse {
	s.Body = v
	return s
}

func (s *ListQualityArchiveTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
