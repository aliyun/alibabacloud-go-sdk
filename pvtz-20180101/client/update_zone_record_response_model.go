// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateZoneRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateZoneRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateZoneRecordResponseBody) *UpdateZoneRecordResponse
	GetBody() *UpdateZoneRecordResponseBody
}

type UpdateZoneRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZoneRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateZoneRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateZoneRecordResponse) GetBody() *UpdateZoneRecordResponseBody {
	return s.Body
}

func (s *UpdateZoneRecordResponse) SetHeaders(v map[string]*string) *UpdateZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRecordResponse) SetStatusCode(v int32) *UpdateZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZoneRecordResponse) SetBody(v *UpdateZoneRecordResponseBody) *UpdateZoneRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateZoneRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
