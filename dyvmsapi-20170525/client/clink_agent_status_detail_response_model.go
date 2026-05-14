// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkAgentStatusDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkAgentStatusDetailResponse
	GetStatusCode() *int32
	SetBody(v *ClinkAgentStatusDetailResponseBody) *ClinkAgentStatusDetailResponse
	GetBody() *ClinkAgentStatusDetailResponseBody
}

type ClinkAgentStatusDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkAgentStatusDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkAgentStatusDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusDetailResponse) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkAgentStatusDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkAgentStatusDetailResponse) GetBody() *ClinkAgentStatusDetailResponseBody {
	return s.Body
}

func (s *ClinkAgentStatusDetailResponse) SetHeaders(v map[string]*string) *ClinkAgentStatusDetailResponse {
	s.Headers = v
	return s
}

func (s *ClinkAgentStatusDetailResponse) SetStatusCode(v int32) *ClinkAgentStatusDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkAgentStatusDetailResponse) SetBody(v *ClinkAgentStatusDetailResponseBody) *ClinkAgentStatusDetailResponse {
	s.Body = v
	return s
}

func (s *ClinkAgentStatusDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
