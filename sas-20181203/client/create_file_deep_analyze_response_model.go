// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDeepAnalyzeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileDeepAnalyzeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileDeepAnalyzeResponse
	GetStatusCode() *int32
	SetId(v string) *CreateFileDeepAnalyzeResponse
	GetId() *string
	SetEvent(v string) *CreateFileDeepAnalyzeResponse
	GetEvent() *string
	SetBody(v *CreateFileDeepAnalyzeResponseBody) *CreateFileDeepAnalyzeResponse
	GetBody() *CreateFileDeepAnalyzeResponseBody
}

type CreateFileDeepAnalyzeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *CreateFileDeepAnalyzeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileDeepAnalyzeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDeepAnalyzeResponse) GoString() string {
	return s.String()
}

func (s *CreateFileDeepAnalyzeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileDeepAnalyzeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileDeepAnalyzeResponse) GetId() *string {
	return s.Id
}

func (s *CreateFileDeepAnalyzeResponse) GetEvent() *string {
	return s.Event
}

func (s *CreateFileDeepAnalyzeResponse) GetBody() *CreateFileDeepAnalyzeResponseBody {
	return s.Body
}

func (s *CreateFileDeepAnalyzeResponse) SetHeaders(v map[string]*string) *CreateFileDeepAnalyzeResponse {
	s.Headers = v
	return s
}

func (s *CreateFileDeepAnalyzeResponse) SetStatusCode(v int32) *CreateFileDeepAnalyzeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponse) SetId(v string) *CreateFileDeepAnalyzeResponse {
	s.Id = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponse) SetEvent(v string) *CreateFileDeepAnalyzeResponse {
	s.Event = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponse) SetBody(v *CreateFileDeepAnalyzeResponseBody) *CreateFileDeepAnalyzeResponse {
	s.Body = v
	return s
}

func (s *CreateFileDeepAnalyzeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
