// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptHotWordListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryScriptHotWordListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryScriptHotWordListResponse
	GetStatusCode() *int32
	SetBody(v *QueryScriptHotWordListResponseBody) *QueryScriptHotWordListResponse
	GetBody() *QueryScriptHotWordListResponseBody
}

type QueryScriptHotWordListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScriptHotWordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScriptHotWordListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptHotWordListResponse) GoString() string {
	return s.String()
}

func (s *QueryScriptHotWordListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryScriptHotWordListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryScriptHotWordListResponse) GetBody() *QueryScriptHotWordListResponseBody {
	return s.Body
}

func (s *QueryScriptHotWordListResponse) SetHeaders(v map[string]*string) *QueryScriptHotWordListResponse {
	s.Headers = v
	return s
}

func (s *QueryScriptHotWordListResponse) SetStatusCode(v int32) *QueryScriptHotWordListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScriptHotWordListResponse) SetBody(v *QueryScriptHotWordListResponseBody) *QueryScriptHotWordListResponse {
	s.Body = v
	return s
}

func (s *QueryScriptHotWordListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
