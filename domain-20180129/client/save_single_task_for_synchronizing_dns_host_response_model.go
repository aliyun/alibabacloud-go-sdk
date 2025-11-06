// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDnsHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDnsHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDnsHostResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForSynchronizingDnsHostResponseBody) *SaveSingleTaskForSynchronizingDnsHostResponse
	GetBody() *SaveSingleTaskForSynchronizingDnsHostResponseBody
}

type SaveSingleTaskForSynchronizingDnsHostResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForSynchronizingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDnsHostResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) GetBody() *SaveSingleTaskForSynchronizingDnsHostResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetBody(v *SaveSingleTaskForSynchronizingDnsHostResponseBody) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
