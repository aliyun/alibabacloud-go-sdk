// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionDetailFileListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommissionDetailFileListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommissionDetailFileListResponse
	GetStatusCode() *int32
	SetBody(v *GetCommissionDetailFileListResponseBody) *GetCommissionDetailFileListResponse
	GetBody() *GetCommissionDetailFileListResponseBody
}

type GetCommissionDetailFileListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommissionDetailFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommissionDetailFileListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionDetailFileListResponse) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommissionDetailFileListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommissionDetailFileListResponse) GetBody() *GetCommissionDetailFileListResponseBody {
	return s.Body
}

func (s *GetCommissionDetailFileListResponse) SetHeaders(v map[string]*string) *GetCommissionDetailFileListResponse {
	s.Headers = v
	return s
}

func (s *GetCommissionDetailFileListResponse) SetStatusCode(v int32) *GetCommissionDetailFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommissionDetailFileListResponse) SetBody(v *GetCommissionDetailFileListResponseBody) *GetCommissionDetailFileListResponse {
	s.Body = v
	return s
}

func (s *GetCommissionDetailFileListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
