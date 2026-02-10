// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeImageResponse
	GetStatusCode() *int32
	SetId(v string) *AnalyzeImageResponse
	GetId() *string
	SetEvent(v string) *AnalyzeImageResponse
	GetEvent() *string
	SetBody(v *AnalyzeImageResponseBody) *AnalyzeImageResponse
	GetBody() *AnalyzeImageResponseBody
}

type AnalyzeImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                   `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                   `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AnalyzeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeImageResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeImageResponse) GetId() *string {
	return s.Id
}

func (s *AnalyzeImageResponse) GetEvent() *string {
	return s.Event
}

func (s *AnalyzeImageResponse) GetBody() *AnalyzeImageResponseBody {
	return s.Body
}

func (s *AnalyzeImageResponse) SetHeaders(v map[string]*string) *AnalyzeImageResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeImageResponse) SetStatusCode(v int32) *AnalyzeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeImageResponse) SetId(v string) *AnalyzeImageResponse {
	s.Id = &v
	return s
}

func (s *AnalyzeImageResponse) SetEvent(v string) *AnalyzeImageResponse {
	s.Event = &v
	return s
}

func (s *AnalyzeImageResponse) SetBody(v *AnalyzeImageResponseBody) *AnalyzeImageResponse {
	s.Body = v
	return s
}

func (s *AnalyzeImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
