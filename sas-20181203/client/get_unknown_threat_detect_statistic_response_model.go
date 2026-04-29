// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnknownThreatDetectStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUnknownThreatDetectStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUnknownThreatDetectStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetUnknownThreatDetectStatisticResponseBody) *GetUnknownThreatDetectStatisticResponse
	GetBody() *GetUnknownThreatDetectStatisticResponseBody
}

type GetUnknownThreatDetectStatisticResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnknownThreatDetectStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnknownThreatDetectStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUnknownThreatDetectStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetUnknownThreatDetectStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUnknownThreatDetectStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnknownThreatDetectStatisticResponse) GetBody() *GetUnknownThreatDetectStatisticResponseBody {
	return s.Body
}

func (s *GetUnknownThreatDetectStatisticResponse) SetHeaders(v map[string]*string) *GetUnknownThreatDetectStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponse) SetStatusCode(v int32) *GetUnknownThreatDetectStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponse) SetBody(v *GetUnknownThreatDetectStatisticResponseBody) *GetUnknownThreatDetectStatisticResponse {
	s.Body = v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
