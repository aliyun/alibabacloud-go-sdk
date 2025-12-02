// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKeywordLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKeywordLibResponse
	GetStatusCode() *int32
	SetBody(v *GetKeywordLibResponseBody) *GetKeywordLibResponse
	GetBody() *GetKeywordLibResponseBody
}

type GetKeywordLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordLibResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKeywordLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKeywordLibResponse) GetBody() *GetKeywordLibResponseBody {
	return s.Body
}

func (s *GetKeywordLibResponse) SetHeaders(v map[string]*string) *GetKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordLibResponse) SetStatusCode(v int32) *GetKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordLibResponse) SetBody(v *GetKeywordLibResponseBody) *GetKeywordLibResponse {
	s.Body = v
	return s
}

func (s *GetKeywordLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
