// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorkItemFeedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorkItemFeedResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorkItemFeedResponseBody) *QueryWorkItemFeedResponse
	GetBody() *QueryWorkItemFeedResponseBody
}

type QueryWorkItemFeedResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorkItemFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorkItemFeedResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponse) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorkItemFeedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorkItemFeedResponse) GetBody() *QueryWorkItemFeedResponseBody {
	return s.Body
}

func (s *QueryWorkItemFeedResponse) SetHeaders(v map[string]*string) *QueryWorkItemFeedResponse {
	s.Headers = v
	return s
}

func (s *QueryWorkItemFeedResponse) SetStatusCode(v int32) *QueryWorkItemFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorkItemFeedResponse) SetBody(v *QueryWorkItemFeedResponseBody) *QueryWorkItemFeedResponse {
	s.Body = v
	return s
}

func (s *QueryWorkItemFeedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
