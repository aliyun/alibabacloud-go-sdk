// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublishBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPublishBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPublishBatchResponse
	GetStatusCode() *int32
	SetBody(v *AddPublishBatchResponseBody) *AddPublishBatchResponse
	GetBody() *AddPublishBatchResponseBody
}

type AddPublishBatchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPublishBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPublishBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPublishBatchResponse) GoString() string {
	return s.String()
}

func (s *AddPublishBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPublishBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPublishBatchResponse) GetBody() *AddPublishBatchResponseBody {
	return s.Body
}

func (s *AddPublishBatchResponse) SetHeaders(v map[string]*string) *AddPublishBatchResponse {
	s.Headers = v
	return s
}

func (s *AddPublishBatchResponse) SetStatusCode(v int32) *AddPublishBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPublishBatchResponse) SetBody(v *AddPublishBatchResponseBody) *AddPublishBatchResponse {
	s.Body = v
	return s
}

func (s *AddPublishBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
