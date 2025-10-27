// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublishBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublishBatchResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublishBatchResponseBody) *UpdatePublishBatchResponse
	GetBody() *UpdatePublishBatchResponseBody
}

type UpdatePublishBatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublishBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublishBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublishBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublishBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublishBatchResponse) GetBody() *UpdatePublishBatchResponseBody {
	return s.Body
}

func (s *UpdatePublishBatchResponse) SetHeaders(v map[string]*string) *UpdatePublishBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublishBatchResponse) SetStatusCode(v int32) *UpdatePublishBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublishBatchResponse) SetBody(v *UpdatePublishBatchResponseBody) *UpdatePublishBatchResponse {
	s.Body = v
	return s
}

func (s *UpdatePublishBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
