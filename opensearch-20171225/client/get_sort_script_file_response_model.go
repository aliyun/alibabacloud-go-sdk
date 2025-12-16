// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSortScriptFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSortScriptFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSortScriptFileResponse
	GetStatusCode() *int32
	SetBody(v *GetSortScriptFileResponseBody) *GetSortScriptFileResponse
	GetBody() *GetSortScriptFileResponseBody
}

type GetSortScriptFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSortScriptFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSortScriptFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSortScriptFileResponse) GetBody() *GetSortScriptFileResponseBody {
	return s.Body
}

func (s *GetSortScriptFileResponse) SetHeaders(v map[string]*string) *GetSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptFileResponse) SetStatusCode(v int32) *GetSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSortScriptFileResponse) SetBody(v *GetSortScriptFileResponseBody) *GetSortScriptFileResponse {
	s.Body = v
	return s
}

func (s *GetSortScriptFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
