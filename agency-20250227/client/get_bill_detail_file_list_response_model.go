// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillDetailFileListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBillDetailFileListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBillDetailFileListResponse
	GetStatusCode() *int32
	SetBody(v *GetBillDetailFileListResponseBody) *GetBillDetailFileListResponse
	GetBody() *GetBillDetailFileListResponseBody
}

type GetBillDetailFileListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBillDetailFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBillDetailFileListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBillDetailFileListResponse) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBillDetailFileListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBillDetailFileListResponse) GetBody() *GetBillDetailFileListResponseBody {
	return s.Body
}

func (s *GetBillDetailFileListResponse) SetHeaders(v map[string]*string) *GetBillDetailFileListResponse {
	s.Headers = v
	return s
}

func (s *GetBillDetailFileListResponse) SetStatusCode(v int32) *GetBillDetailFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBillDetailFileListResponse) SetBody(v *GetBillDetailFileListResponseBody) *GetBillDetailFileListResponse {
	s.Body = v
	return s
}

func (s *GetBillDetailFileListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
