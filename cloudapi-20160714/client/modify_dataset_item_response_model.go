// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatasetItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatasetItemResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatasetItemResponseBody) *ModifyDatasetItemResponse
	GetBody() *ModifyDatasetItemResponseBody
}

type ModifyDatasetItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatasetItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatasetItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatasetItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatasetItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatasetItemResponse) GetBody() *ModifyDatasetItemResponseBody {
	return s.Body
}

func (s *ModifyDatasetItemResponse) SetHeaders(v map[string]*string) *ModifyDatasetItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatasetItemResponse) SetStatusCode(v int32) *ModifyDatasetItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatasetItemResponse) SetBody(v *ModifyDatasetItemResponseBody) *ModifyDatasetItemResponse {
	s.Body = v
	return s
}

func (s *ModifyDatasetItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
