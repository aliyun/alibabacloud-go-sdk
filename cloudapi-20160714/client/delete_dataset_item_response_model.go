// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetItemResponseBody) *DeleteDatasetItemResponse
	GetBody() *DeleteDatasetItemResponseBody
}

type DeleteDatasetItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetItemResponse) GetBody() *DeleteDatasetItemResponseBody {
	return s.Body
}

func (s *DeleteDatasetItemResponse) SetHeaders(v map[string]*string) *DeleteDatasetItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetItemResponse) SetStatusCode(v int32) *DeleteDatasetItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetItemResponse) SetBody(v *DeleteDatasetItemResponseBody) *DeleteDatasetItemResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetItemResponse) Validate() error {
	return dara.Validate(s)
}
