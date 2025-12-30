// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteZoneRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteZoneRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteZoneRecordResponseBody) *DeleteZoneRecordResponse
	GetBody() *DeleteZoneRecordResponseBody
}

type DeleteZoneRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteZoneRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteZoneRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteZoneRecordResponse) GetBody() *DeleteZoneRecordResponseBody {
	return s.Body
}

func (s *DeleteZoneRecordResponse) SetHeaders(v map[string]*string) *DeleteZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneRecordResponse) SetStatusCode(v int32) *DeleteZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteZoneRecordResponse) SetBody(v *DeleteZoneRecordResponseBody) *DeleteZoneRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteZoneRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
