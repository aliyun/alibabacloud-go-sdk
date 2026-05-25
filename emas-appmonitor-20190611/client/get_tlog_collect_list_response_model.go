// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogCollectListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlogCollectListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlogCollectListResponse
	GetStatusCode() *int32
	SetBody(v *GetTlogCollectListResponseBody) *GetTlogCollectListResponse
	GetBody() *GetTlogCollectListResponseBody
}

type GetTlogCollectListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlogCollectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlogCollectListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlogCollectListResponse) GoString() string {
	return s.String()
}

func (s *GetTlogCollectListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlogCollectListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlogCollectListResponse) GetBody() *GetTlogCollectListResponseBody {
	return s.Body
}

func (s *GetTlogCollectListResponse) SetHeaders(v map[string]*string) *GetTlogCollectListResponse {
	s.Headers = v
	return s
}

func (s *GetTlogCollectListResponse) SetStatusCode(v int32) *GetTlogCollectListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlogCollectListResponse) SetBody(v *GetTlogCollectListResponseBody) *GetTlogCollectListResponse {
	s.Body = v
	return s
}

func (s *GetTlogCollectListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
