// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTargetListByBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTargetListByBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTargetListByBatchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTargetListByBatchResponseBody) *UpdateTargetListByBatchResponse
	GetBody() *UpdateTargetListByBatchResponseBody
}

type UpdateTargetListByBatchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTargetListByBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTargetListByBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTargetListByBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdateTargetListByBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTargetListByBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTargetListByBatchResponse) GetBody() *UpdateTargetListByBatchResponseBody {
	return s.Body
}

func (s *UpdateTargetListByBatchResponse) SetHeaders(v map[string]*string) *UpdateTargetListByBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdateTargetListByBatchResponse) SetStatusCode(v int32) *UpdateTargetListByBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTargetListByBatchResponse) SetBody(v *UpdateTargetListByBatchResponseBody) *UpdateTargetListByBatchResponse {
	s.Body = v
	return s
}

func (s *UpdateTargetListByBatchResponse) Validate() error {
	return dara.Validate(s)
}
