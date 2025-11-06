// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDnsHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDnsHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForModifyingDnsHostResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForModifyingDnsHostResponseBody) *SaveSingleTaskForModifyingDnsHostResponse
	GetBody() *SaveSingleTaskForModifyingDnsHostResponseBody
}

type SaveSingleTaskForModifyingDnsHostResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForModifyingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForModifyingDnsHostResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) GetBody() *SaveSingleTaskForModifyingDnsHostResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForModifyingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetBody(v *SaveSingleTaskForModifyingDnsHostResponseBody) *SaveSingleTaskForModifyingDnsHostResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
