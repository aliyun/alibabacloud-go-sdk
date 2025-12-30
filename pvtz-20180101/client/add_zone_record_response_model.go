// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddZoneRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddZoneRecordResponse
	GetStatusCode() *int32
	SetBody(v *AddZoneRecordResponseBody) *AddZoneRecordResponse
	GetBody() *AddZoneRecordResponseBody
}

type AddZoneRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZoneRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddZoneRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddZoneRecordResponse) GetBody() *AddZoneRecordResponseBody {
	return s.Body
}

func (s *AddZoneRecordResponse) SetHeaders(v map[string]*string) *AddZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *AddZoneRecordResponse) SetStatusCode(v int32) *AddZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneRecordResponse) SetBody(v *AddZoneRecordResponseBody) *AddZoneRecordResponse {
	s.Body = v
	return s
}

func (s *AddZoneRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
