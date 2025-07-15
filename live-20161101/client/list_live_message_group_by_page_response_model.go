// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveMessageGroupByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveMessageGroupByPageResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveMessageGroupByPageResponseBody) *ListLiveMessageGroupByPageResponse
	GetBody() *ListLiveMessageGroupByPageResponseBody
}

type ListLiveMessageGroupByPageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveMessageGroupByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveMessageGroupByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupByPageResponse) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveMessageGroupByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveMessageGroupByPageResponse) GetBody() *ListLiveMessageGroupByPageResponseBody {
	return s.Body
}

func (s *ListLiveMessageGroupByPageResponse) SetHeaders(v map[string]*string) *ListLiveMessageGroupByPageResponse {
	s.Headers = v
	return s
}

func (s *ListLiveMessageGroupByPageResponse) SetStatusCode(v int32) *ListLiveMessageGroupByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponse) SetBody(v *ListLiveMessageGroupByPageResponseBody) *ListLiveMessageGroupByPageResponse {
	s.Body = v
	return s
}

func (s *ListLiveMessageGroupByPageResponse) Validate() error {
	return dara.Validate(s)
}
