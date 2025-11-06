// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDnsHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDnsHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForDeletingDnsHostResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForDeletingDnsHostResponseBody) *SaveSingleTaskForDeletingDnsHostResponse
	GetBody() *SaveSingleTaskForDeletingDnsHostResponseBody
}

type SaveSingleTaskForDeletingDnsHostResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForDeletingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForDeletingDnsHostResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) GetBody() *SaveSingleTaskForDeletingDnsHostResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForDeletingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetBody(v *SaveSingleTaskForDeletingDnsHostResponseBody) *SaveSingleTaskForDeletingDnsHostResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
