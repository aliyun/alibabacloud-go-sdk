// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSortScriptFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSortScriptFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSortScriptFileResponse
	GetStatusCode() *int32
	SetBody(v *SaveSortScriptFileResponseBody) *SaveSortScriptFileResponse
	GetBody() *SaveSortScriptFileResponseBody
}

type SaveSortScriptFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSortScriptFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSortScriptFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSortScriptFileResponse) GetBody() *SaveSortScriptFileResponseBody {
	return s.Body
}

func (s *SaveSortScriptFileResponse) SetHeaders(v map[string]*string) *SaveSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *SaveSortScriptFileResponse) SetStatusCode(v int32) *SaveSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSortScriptFileResponse) SetBody(v *SaveSortScriptFileResponseBody) *SaveSortScriptFileResponse {
	s.Body = v
	return s
}

func (s *SaveSortScriptFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
