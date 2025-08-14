// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaLiveChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaLiveChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaLiveChannelsResponseBody) *ListMediaLiveChannelsResponse
	GetBody() *ListMediaLiveChannelsResponseBody
}

type ListMediaLiveChannelsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaLiveChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaLiveChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaLiveChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaLiveChannelsResponse) GetBody() *ListMediaLiveChannelsResponseBody {
	return s.Body
}

func (s *ListMediaLiveChannelsResponse) SetHeaders(v map[string]*string) *ListMediaLiveChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaLiveChannelsResponse) SetStatusCode(v int32) *ListMediaLiveChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaLiveChannelsResponse) SetBody(v *ListMediaLiveChannelsResponseBody) *ListMediaLiveChannelsResponse {
	s.Body = v
	return s
}

func (s *ListMediaLiveChannelsResponse) Validate() error {
	return dara.Validate(s)
}
