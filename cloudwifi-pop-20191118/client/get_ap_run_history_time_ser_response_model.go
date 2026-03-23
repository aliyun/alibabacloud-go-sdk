// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApRunHistoryTimeSerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApRunHistoryTimeSerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApRunHistoryTimeSerResponse
	GetStatusCode() *int32
	SetBody(v *GetApRunHistoryTimeSerResponseBody) *GetApRunHistoryTimeSerResponse
	GetBody() *GetApRunHistoryTimeSerResponseBody
}

type GetApRunHistoryTimeSerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApRunHistoryTimeSerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApRunHistoryTimeSerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApRunHistoryTimeSerResponse) GoString() string {
	return s.String()
}

func (s *GetApRunHistoryTimeSerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApRunHistoryTimeSerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApRunHistoryTimeSerResponse) GetBody() *GetApRunHistoryTimeSerResponseBody {
	return s.Body
}

func (s *GetApRunHistoryTimeSerResponse) SetHeaders(v map[string]*string) *GetApRunHistoryTimeSerResponse {
	s.Headers = v
	return s
}

func (s *GetApRunHistoryTimeSerResponse) SetStatusCode(v int32) *GetApRunHistoryTimeSerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApRunHistoryTimeSerResponse) SetBody(v *GetApRunHistoryTimeSerResponseBody) *GetApRunHistoryTimeSerResponse {
	s.Body = v
	return s
}

func (s *GetApRunHistoryTimeSerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
