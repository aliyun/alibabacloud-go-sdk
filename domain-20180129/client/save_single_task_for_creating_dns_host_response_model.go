// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingDnsHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingDnsHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCreatingDnsHostResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCreatingDnsHostResponseBody) *SaveSingleTaskForCreatingDnsHostResponse
	GetBody() *SaveSingleTaskForCreatingDnsHostResponseBody
}

type SaveSingleTaskForCreatingDnsHostResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCreatingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCreatingDnsHostResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) GetBody() *SaveSingleTaskForCreatingDnsHostResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetBody(v *SaveSingleTaskForCreatingDnsHostResponseBody) *SaveSingleTaskForCreatingDnsHostResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
