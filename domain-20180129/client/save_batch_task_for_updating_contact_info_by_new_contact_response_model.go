// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByNewContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse
	GetBody() *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody
}

type SaveBatchTaskForUpdatingContactInfoByNewContactResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) GetBody() *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetBody(v *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
