// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCnnfStatusDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterCnnfStatusDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterCnnfStatusDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterCnnfStatusDetailResponseBody) *ListClusterCnnfStatusDetailResponse
	GetBody() *ListClusterCnnfStatusDetailResponseBody
}

type ListClusterCnnfStatusDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterCnnfStatusDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterCnnfStatusDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCnnfStatusDetailResponse) GoString() string {
	return s.String()
}

func (s *ListClusterCnnfStatusDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterCnnfStatusDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterCnnfStatusDetailResponse) GetBody() *ListClusterCnnfStatusDetailResponseBody {
	return s.Body
}

func (s *ListClusterCnnfStatusDetailResponse) SetHeaders(v map[string]*string) *ListClusterCnnfStatusDetailResponse {
	s.Headers = v
	return s
}

func (s *ListClusterCnnfStatusDetailResponse) SetStatusCode(v int32) *ListClusterCnnfStatusDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponse) SetBody(v *ListClusterCnnfStatusDetailResponseBody) *ListClusterCnnfStatusDetailResponse {
	s.Body = v
	return s
}

func (s *ListClusterCnnfStatusDetailResponse) Validate() error {
	return dara.Validate(s)
}
