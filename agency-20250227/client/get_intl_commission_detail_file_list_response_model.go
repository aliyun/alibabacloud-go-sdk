// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlCommissionDetailFileListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntlCommissionDetailFileListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntlCommissionDetailFileListResponse
	GetStatusCode() *int32
	SetBody(v *GetIntlCommissionDetailFileListResponseBody) *GetIntlCommissionDetailFileListResponse
	GetBody() *GetIntlCommissionDetailFileListResponseBody
}

type GetIntlCommissionDetailFileListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntlCommissionDetailFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntlCommissionDetailFileListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntlCommissionDetailFileListResponse) GoString() string {
	return s.String()
}

func (s *GetIntlCommissionDetailFileListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntlCommissionDetailFileListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntlCommissionDetailFileListResponse) GetBody() *GetIntlCommissionDetailFileListResponseBody {
	return s.Body
}

func (s *GetIntlCommissionDetailFileListResponse) SetHeaders(v map[string]*string) *GetIntlCommissionDetailFileListResponse {
	s.Headers = v
	return s
}

func (s *GetIntlCommissionDetailFileListResponse) SetStatusCode(v int32) *GetIntlCommissionDetailFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntlCommissionDetailFileListResponse) SetBody(v *GetIntlCommissionDetailFileListResponseBody) *GetIntlCommissionDetailFileListResponse {
	s.Body = v
	return s
}

func (s *GetIntlCommissionDetailFileListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
