// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForUpdatingContactInfoResponseBody) *SaveBatchTaskForUpdatingContactInfoResponse
	GetBody() *SaveBatchTaskForUpdatingContactInfoResponseBody
}

type SaveBatchTaskForUpdatingContactInfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForUpdatingContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) GetBody() *SaveBatchTaskForUpdatingContactInfoResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) SetBody(v *SaveBatchTaskForUpdatingContactInfoResponseBody) *SaveBatchTaskForUpdatingContactInfoResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
