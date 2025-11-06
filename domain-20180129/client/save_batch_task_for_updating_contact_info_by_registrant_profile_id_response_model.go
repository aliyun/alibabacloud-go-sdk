// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse
	GetBody() *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse struct {
	Headers    map[string]*string                                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) GetBody() *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetBody(v *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
