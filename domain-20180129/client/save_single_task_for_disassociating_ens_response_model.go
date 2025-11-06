// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDisassociatingEnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForDisassociatingEnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForDisassociatingEnsResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForDisassociatingEnsResponseBody) *SaveSingleTaskForDisassociatingEnsResponse
	GetBody() *SaveSingleTaskForDisassociatingEnsResponseBody
}

type SaveSingleTaskForDisassociatingEnsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForDisassociatingEnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForDisassociatingEnsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) GetBody() *SaveSingleTaskForDisassociatingEnsResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDisassociatingEnsResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetStatusCode(v int32) *SaveSingleTaskForDisassociatingEnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetBody(v *SaveSingleTaskForDisassociatingEnsResponseBody) *SaveSingleTaskForDisassociatingEnsResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
