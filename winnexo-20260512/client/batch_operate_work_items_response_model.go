// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateWorkItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchOperateWorkItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchOperateWorkItemsResponse
	GetStatusCode() *int32
	SetBody(v *BatchOperateWorkItemsResponseBody) *BatchOperateWorkItemsResponse
	GetBody() *BatchOperateWorkItemsResponseBody
}

type BatchOperateWorkItemsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchOperateWorkItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchOperateWorkItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateWorkItemsResponse) GoString() string {
	return s.String()
}

func (s *BatchOperateWorkItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchOperateWorkItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchOperateWorkItemsResponse) GetBody() *BatchOperateWorkItemsResponseBody {
	return s.Body
}

func (s *BatchOperateWorkItemsResponse) SetHeaders(v map[string]*string) *BatchOperateWorkItemsResponse {
	s.Headers = v
	return s
}

func (s *BatchOperateWorkItemsResponse) SetStatusCode(v int32) *BatchOperateWorkItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchOperateWorkItemsResponse) SetBody(v *BatchOperateWorkItemsResponseBody) *BatchOperateWorkItemsResponse {
	s.Body = v
	return s
}

func (s *BatchOperateWorkItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
