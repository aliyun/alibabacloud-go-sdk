// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableIntroWikiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableIntroWikiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableIntroWikiResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableIntroWikiResponseBody) *GetMetaTableIntroWikiResponse
	GetBody() *GetMetaTableIntroWikiResponseBody
}

type GetMetaTableIntroWikiResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableIntroWikiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableIntroWikiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableIntroWikiResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableIntroWikiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableIntroWikiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableIntroWikiResponse) GetBody() *GetMetaTableIntroWikiResponseBody {
	return s.Body
}

func (s *GetMetaTableIntroWikiResponse) SetHeaders(v map[string]*string) *GetMetaTableIntroWikiResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableIntroWikiResponse) SetStatusCode(v int32) *GetMetaTableIntroWikiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableIntroWikiResponse) SetBody(v *GetMetaTableIntroWikiResponseBody) *GetMetaTableIntroWikiResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableIntroWikiResponse) Validate() error {
	return dara.Validate(s)
}
