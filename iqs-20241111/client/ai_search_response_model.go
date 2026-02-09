// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AiSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AiSearchResponse
	GetStatusCode() *int32
	SetId(v string) *AiSearchResponse
	GetId() *string
	SetEvent(v string) *AiSearchResponse
	GetEvent() *string
	SetBody(v *AiSearchResponseBody) *AiSearchResponse
	GetBody() *AiSearchResponseBody
}

type AiSearchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string               `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string               `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AiSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s AiSearchResponse) GoString() string {
	return s.String()
}

func (s *AiSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AiSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AiSearchResponse) GetId() *string {
	return s.Id
}

func (s *AiSearchResponse) GetEvent() *string {
	return s.Event
}

func (s *AiSearchResponse) GetBody() *AiSearchResponseBody {
	return s.Body
}

func (s *AiSearchResponse) SetHeaders(v map[string]*string) *AiSearchResponse {
	s.Headers = v
	return s
}

func (s *AiSearchResponse) SetStatusCode(v int32) *AiSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AiSearchResponse) SetId(v string) *AiSearchResponse {
	s.Id = &v
	return s
}

func (s *AiSearchResponse) SetEvent(v string) *AiSearchResponse {
	s.Event = &v
	return s
}

func (s *AiSearchResponse) SetBody(v *AiSearchResponseBody) *AiSearchResponse {
	s.Body = v
	return s
}

func (s *AiSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
