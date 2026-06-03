// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdatingContactInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForUpdatingContactInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForUpdatingContactInfoResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForUpdatingContactInfoResponseBody) *SaveSingleTaskForUpdatingContactInfoResponse
	GetBody() *SaveSingleTaskForUpdatingContactInfoResponseBody
}

type SaveSingleTaskForUpdatingContactInfoResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForUpdatingContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForUpdatingContactInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) GetBody() *SaveSingleTaskForUpdatingContactInfoResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetStatusCode(v int32) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetBody(v *SaveSingleTaskForUpdatingContactInfoResponseBody) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
