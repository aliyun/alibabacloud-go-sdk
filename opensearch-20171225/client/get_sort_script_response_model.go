// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSortScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSortScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSortScriptResponse
	GetStatusCode() *int32
	SetBody(v *GetSortScriptResponseBody) *GetSortScriptResponse
	GetBody() *GetSortScriptResponseBody
}

type GetSortScriptResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSortScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSortScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSortScriptResponse) GetBody() *GetSortScriptResponseBody {
	return s.Body
}

func (s *GetSortScriptResponse) SetHeaders(v map[string]*string) *GetSortScriptResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptResponse) SetStatusCode(v int32) *GetSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSortScriptResponse) SetBody(v *GetSortScriptResponseBody) *GetSortScriptResponse {
	s.Body = v
	return s
}

func (s *GetSortScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
