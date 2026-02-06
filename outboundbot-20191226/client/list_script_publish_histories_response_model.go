// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptPublishHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptPublishHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptPublishHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptPublishHistoriesResponseBody) *ListScriptPublishHistoriesResponse
	GetBody() *ListScriptPublishHistoriesResponseBody
}

type ListScriptPublishHistoriesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptPublishHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptPublishHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptPublishHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListScriptPublishHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptPublishHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptPublishHistoriesResponse) GetBody() *ListScriptPublishHistoriesResponseBody {
	return s.Body
}

func (s *ListScriptPublishHistoriesResponse) SetHeaders(v map[string]*string) *ListScriptPublishHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListScriptPublishHistoriesResponse) SetStatusCode(v int32) *ListScriptPublishHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptPublishHistoriesResponse) SetBody(v *ListScriptPublishHistoriesResponseBody) *ListScriptPublishHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListScriptPublishHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
