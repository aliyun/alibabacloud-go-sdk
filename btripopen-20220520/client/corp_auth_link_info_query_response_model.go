// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpAuthLinkInfoQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CorpAuthLinkInfoQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CorpAuthLinkInfoQueryResponse
	GetStatusCode() *int32
	SetBody(v *CorpAuthLinkInfoQueryResponseBody) *CorpAuthLinkInfoQueryResponse
	GetBody() *CorpAuthLinkInfoQueryResponseBody
}

type CorpAuthLinkInfoQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CorpAuthLinkInfoQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CorpAuthLinkInfoQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CorpAuthLinkInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *CorpAuthLinkInfoQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CorpAuthLinkInfoQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CorpAuthLinkInfoQueryResponse) GetBody() *CorpAuthLinkInfoQueryResponseBody {
	return s.Body
}

func (s *CorpAuthLinkInfoQueryResponse) SetHeaders(v map[string]*string) *CorpAuthLinkInfoQueryResponse {
	s.Headers = v
	return s
}

func (s *CorpAuthLinkInfoQueryResponse) SetStatusCode(v int32) *CorpAuthLinkInfoQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponse) SetBody(v *CorpAuthLinkInfoQueryResponseBody) *CorpAuthLinkInfoQueryResponse {
	s.Body = v
	return s
}

func (s *CorpAuthLinkInfoQueryResponse) Validate() error {
	return dara.Validate(s)
}
