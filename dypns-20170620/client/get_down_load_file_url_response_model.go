// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownLoadFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDownLoadFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDownLoadFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetDownLoadFileUrlResponseBody) *GetDownLoadFileUrlResponse
	GetBody() *GetDownLoadFileUrlResponseBody
}

type GetDownLoadFileUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDownLoadFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDownLoadFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDownLoadFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDownLoadFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDownLoadFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDownLoadFileUrlResponse) GetBody() *GetDownLoadFileUrlResponseBody {
	return s.Body
}

func (s *GetDownLoadFileUrlResponse) SetHeaders(v map[string]*string) *GetDownLoadFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDownLoadFileUrlResponse) SetStatusCode(v int32) *GetDownLoadFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDownLoadFileUrlResponse) SetBody(v *GetDownLoadFileUrlResponseBody) *GetDownLoadFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetDownLoadFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
