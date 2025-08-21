// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateFirstFloorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommonCateFirstFloorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommonCateFirstFloorResponse
	GetStatusCode() *int32
	SetBody(v *ListCommonCateFirstFloorResponseBody) *ListCommonCateFirstFloorResponse
	GetBody() *ListCommonCateFirstFloorResponseBody
}

type ListCommonCateFirstFloorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonCateFirstFloorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonCateFirstFloorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateFirstFloorResponse) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommonCateFirstFloorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommonCateFirstFloorResponse) GetBody() *ListCommonCateFirstFloorResponseBody {
	return s.Body
}

func (s *ListCommonCateFirstFloorResponse) SetHeaders(v map[string]*string) *ListCommonCateFirstFloorResponse {
	s.Headers = v
	return s
}

func (s *ListCommonCateFirstFloorResponse) SetStatusCode(v int32) *ListCommonCateFirstFloorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonCateFirstFloorResponse) SetBody(v *ListCommonCateFirstFloorResponseBody) *ListCommonCateFirstFloorResponse {
	s.Body = v
	return s
}

func (s *ListCommonCateFirstFloorResponse) Validate() error {
	return dara.Validate(s)
}
