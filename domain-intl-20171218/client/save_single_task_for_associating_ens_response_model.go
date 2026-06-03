// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAssociatingEnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForAssociatingEnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForAssociatingEnsResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForAssociatingEnsResponseBody) *SaveSingleTaskForAssociatingEnsResponse
	GetBody() *SaveSingleTaskForAssociatingEnsResponseBody
}

type SaveSingleTaskForAssociatingEnsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForAssociatingEnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForAssociatingEnsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForAssociatingEnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForAssociatingEnsResponse) GetBody() *SaveSingleTaskForAssociatingEnsResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForAssociatingEnsResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetStatusCode(v int32) *SaveSingleTaskForAssociatingEnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetBody(v *SaveSingleTaskForAssociatingEnsResponseBody) *SaveSingleTaskForAssociatingEnsResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
