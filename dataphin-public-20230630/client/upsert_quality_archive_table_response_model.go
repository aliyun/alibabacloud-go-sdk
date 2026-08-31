// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityArchiveTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityArchiveTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityArchiveTableResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityArchiveTableResponseBody) *UpsertQualityArchiveTableResponse
	GetBody() *UpsertQualityArchiveTableResponseBody
}

type UpsertQualityArchiveTableResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityArchiveTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityArchiveTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityArchiveTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityArchiveTableResponse) GetBody() *UpsertQualityArchiveTableResponseBody {
	return s.Body
}

func (s *UpsertQualityArchiveTableResponse) SetHeaders(v map[string]*string) *UpsertQualityArchiveTableResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityArchiveTableResponse) SetStatusCode(v int32) *UpsertQualityArchiveTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityArchiveTableResponse) SetBody(v *UpsertQualityArchiveTableResponseBody) *UpsertQualityArchiveTableResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityArchiveTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
