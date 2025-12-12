// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncForLLMResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDataSyncForLLMResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDataSyncForLLMResponse
	GetStatusCode() *int32
	SetBody(v *UploadDataSyncForLLMResponseBody) *UploadDataSyncForLLMResponse
	GetBody() *UploadDataSyncForLLMResponseBody
}

type UploadDataSyncForLLMResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataSyncForLLMResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataSyncForLLMResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMResponse) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDataSyncForLLMResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDataSyncForLLMResponse) GetBody() *UploadDataSyncForLLMResponseBody {
	return s.Body
}

func (s *UploadDataSyncForLLMResponse) SetHeaders(v map[string]*string) *UploadDataSyncForLLMResponse {
	s.Headers = v
	return s
}

func (s *UploadDataSyncForLLMResponse) SetStatusCode(v int32) *UploadDataSyncForLLMResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataSyncForLLMResponse) SetBody(v *UploadDataSyncForLLMResponseBody) *UploadDataSyncForLLMResponse {
	s.Body = v
	return s
}

func (s *UploadDataSyncForLLMResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
