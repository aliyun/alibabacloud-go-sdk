// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableIntroWikiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaTableIntroWikiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaTableIntroWikiResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaTableIntroWikiResponseBody) *UpdateMetaTableIntroWikiResponse
	GetBody() *UpdateMetaTableIntroWikiResponseBody
}

type UpdateMetaTableIntroWikiResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaTableIntroWikiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaTableIntroWikiResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableIntroWikiResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableIntroWikiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaTableIntroWikiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaTableIntroWikiResponse) GetBody() *UpdateMetaTableIntroWikiResponseBody {
	return s.Body
}

func (s *UpdateMetaTableIntroWikiResponse) SetHeaders(v map[string]*string) *UpdateMetaTableIntroWikiResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaTableIntroWikiResponse) SetStatusCode(v int32) *UpdateMetaTableIntroWikiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaTableIntroWikiResponse) SetBody(v *UpdateMetaTableIntroWikiResponseBody) *UpdateMetaTableIntroWikiResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaTableIntroWikiResponse) Validate() error {
	return dara.Validate(s)
}
