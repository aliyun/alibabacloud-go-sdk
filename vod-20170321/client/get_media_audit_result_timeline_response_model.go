// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaAuditResultTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaAuditResultTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaAuditResultTimelineResponseBody) *GetMediaAuditResultTimelineResponse
	GetBody() *GetMediaAuditResultTimelineResponseBody
}

type GetMediaAuditResultTimelineResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaAuditResultTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaAuditResultTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaAuditResultTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaAuditResultTimelineResponse) GetBody() *GetMediaAuditResultTimelineResponseBody {
	return s.Body
}

func (s *GetMediaAuditResultTimelineResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultTimelineResponse) SetStatusCode(v int32) *GetMediaAuditResultTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponse) SetBody(v *GetMediaAuditResultTimelineResponseBody) *GetMediaAuditResultTimelineResponse {
	s.Body = v
	return s
}

func (s *GetMediaAuditResultTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
