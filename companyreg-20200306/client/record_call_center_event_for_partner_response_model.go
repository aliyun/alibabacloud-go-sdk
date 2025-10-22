// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordCallCenterEventForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecordCallCenterEventForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecordCallCenterEventForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *RecordCallCenterEventForPartnerResponseBody) *RecordCallCenterEventForPartnerResponse
	GetBody() *RecordCallCenterEventForPartnerResponseBody
}

type RecordCallCenterEventForPartnerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecordCallCenterEventForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecordCallCenterEventForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s RecordCallCenterEventForPartnerResponse) GoString() string {
	return s.String()
}

func (s *RecordCallCenterEventForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecordCallCenterEventForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecordCallCenterEventForPartnerResponse) GetBody() *RecordCallCenterEventForPartnerResponseBody {
	return s.Body
}

func (s *RecordCallCenterEventForPartnerResponse) SetHeaders(v map[string]*string) *RecordCallCenterEventForPartnerResponse {
	s.Headers = v
	return s
}

func (s *RecordCallCenterEventForPartnerResponse) SetStatusCode(v int32) *RecordCallCenterEventForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponse) SetBody(v *RecordCallCenterEventForPartnerResponseBody) *RecordCallCenterEventForPartnerResponse {
	s.Body = v
	return s
}

func (s *RecordCallCenterEventForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
