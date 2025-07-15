// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpsItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpsItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpsItemsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpsItemsResponseBody) *DeleteOpsItemsResponse
	GetBody() *DeleteOpsItemsResponseBody
}

type DeleteOpsItemsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpsItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpsItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpsItemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpsItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpsItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpsItemsResponse) GetBody() *DeleteOpsItemsResponseBody {
	return s.Body
}

func (s *DeleteOpsItemsResponse) SetHeaders(v map[string]*string) *DeleteOpsItemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpsItemsResponse) SetStatusCode(v int32) *DeleteOpsItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpsItemsResponse) SetBody(v *DeleteOpsItemsResponseBody) *DeleteOpsItemsResponse {
	s.Body = v
	return s
}

func (s *DeleteOpsItemsResponse) Validate() error {
	return dara.Validate(s)
}
