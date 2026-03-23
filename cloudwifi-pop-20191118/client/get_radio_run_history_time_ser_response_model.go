// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRadioRunHistoryTimeSerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRadioRunHistoryTimeSerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRadioRunHistoryTimeSerResponse
	GetStatusCode() *int32
	SetBody(v *GetRadioRunHistoryTimeSerResponseBody) *GetRadioRunHistoryTimeSerResponse
	GetBody() *GetRadioRunHistoryTimeSerResponseBody
}

type GetRadioRunHistoryTimeSerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRadioRunHistoryTimeSerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRadioRunHistoryTimeSerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRadioRunHistoryTimeSerResponse) GoString() string {
	return s.String()
}

func (s *GetRadioRunHistoryTimeSerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRadioRunHistoryTimeSerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRadioRunHistoryTimeSerResponse) GetBody() *GetRadioRunHistoryTimeSerResponseBody {
	return s.Body
}

func (s *GetRadioRunHistoryTimeSerResponse) SetHeaders(v map[string]*string) *GetRadioRunHistoryTimeSerResponse {
	s.Headers = v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponse) SetStatusCode(v int32) *GetRadioRunHistoryTimeSerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponse) SetBody(v *GetRadioRunHistoryTimeSerResponseBody) *GetRadioRunHistoryTimeSerResponse {
	s.Body = v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
