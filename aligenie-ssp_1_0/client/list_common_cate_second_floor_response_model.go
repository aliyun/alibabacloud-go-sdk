// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateSecondFloorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommonCateSecondFloorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommonCateSecondFloorResponse
	GetStatusCode() *int32
	SetBody(v *ListCommonCateSecondFloorResponseBody) *ListCommonCateSecondFloorResponse
	GetBody() *ListCommonCateSecondFloorResponseBody
}

type ListCommonCateSecondFloorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonCateSecondFloorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonCateSecondFloorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateSecondFloorResponse) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommonCateSecondFloorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommonCateSecondFloorResponse) GetBody() *ListCommonCateSecondFloorResponseBody {
	return s.Body
}

func (s *ListCommonCateSecondFloorResponse) SetHeaders(v map[string]*string) *ListCommonCateSecondFloorResponse {
	s.Headers = v
	return s
}

func (s *ListCommonCateSecondFloorResponse) SetStatusCode(v int32) *ListCommonCateSecondFloorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonCateSecondFloorResponse) SetBody(v *ListCommonCateSecondFloorResponseBody) *ListCommonCateSecondFloorResponse {
	s.Body = v
	return s
}

func (s *ListCommonCateSecondFloorResponse) Validate() error {
	return dara.Validate(s)
}
