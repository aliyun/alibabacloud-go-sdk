// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityWatchAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityWatchAlertResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityWatchAlertResponseBody) *UpsertQualityWatchAlertResponse
	GetBody() *UpsertQualityWatchAlertResponseBody
}

type UpsertQualityWatchAlertResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityWatchAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityWatchAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityWatchAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityWatchAlertResponse) GetBody() *UpsertQualityWatchAlertResponseBody {
	return s.Body
}

func (s *UpsertQualityWatchAlertResponse) SetHeaders(v map[string]*string) *UpsertQualityWatchAlertResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityWatchAlertResponse) SetStatusCode(v int32) *UpsertQualityWatchAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityWatchAlertResponse) SetBody(v *UpsertQualityWatchAlertResponseBody) *UpsertQualityWatchAlertResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityWatchAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
