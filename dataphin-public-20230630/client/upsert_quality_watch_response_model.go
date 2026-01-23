// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityWatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityWatchResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityWatchResponseBody) *UpsertQualityWatchResponse
	GetBody() *UpsertQualityWatchResponseBody
}

type UpsertQualityWatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityWatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityWatchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityWatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityWatchResponse) GetBody() *UpsertQualityWatchResponseBody {
	return s.Body
}

func (s *UpsertQualityWatchResponse) SetHeaders(v map[string]*string) *UpsertQualityWatchResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityWatchResponse) SetStatusCode(v int32) *UpsertQualityWatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityWatchResponse) SetBody(v *UpsertQualityWatchResponseBody) *UpsertQualityWatchResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityWatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
