// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocSummaryResponse
	GetStatusCode() *int32
	SetId(v string) *RunDocSummaryResponse
	GetId() *string
	SetEvent(v string) *RunDocSummaryResponse
	GetEvent() *string
	SetBody(v *RunDocSummaryResponseBody) *RunDocSummaryResponse
	GetBody() *RunDocSummaryResponseBody
}

type RunDocSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDocSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponse) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocSummaryResponse) GetId() *string {
	return s.Id
}

func (s *RunDocSummaryResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDocSummaryResponse) GetBody() *RunDocSummaryResponseBody {
	return s.Body
}

func (s *RunDocSummaryResponse) SetHeaders(v map[string]*string) *RunDocSummaryResponse {
	s.Headers = v
	return s
}

func (s *RunDocSummaryResponse) SetStatusCode(v int32) *RunDocSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocSummaryResponse) SetId(v string) *RunDocSummaryResponse {
	s.Id = &v
	return s
}

func (s *RunDocSummaryResponse) SetEvent(v string) *RunDocSummaryResponse {
	s.Event = &v
	return s
}

func (s *RunDocSummaryResponse) SetBody(v *RunDocSummaryResponseBody) *RunDocSummaryResponse {
	s.Body = v
	return s
}

func (s *RunDocSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
