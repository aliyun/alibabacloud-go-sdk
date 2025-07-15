// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetItemResponseBody) *CreateDatasetItemResponse
	GetBody() *CreateDatasetItemResponseBody
}

type CreateDatasetItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetItemResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetItemResponse) GetBody() *CreateDatasetItemResponseBody {
	return s.Body
}

func (s *CreateDatasetItemResponse) SetHeaders(v map[string]*string) *CreateDatasetItemResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetItemResponse) SetStatusCode(v int32) *CreateDatasetItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetItemResponse) SetBody(v *CreateDatasetItemResponseBody) *CreateDatasetItemResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetItemResponse) Validate() error {
	return dara.Validate(s)
}
