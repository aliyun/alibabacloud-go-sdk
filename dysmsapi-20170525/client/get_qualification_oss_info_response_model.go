// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationOssInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualificationOssInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualificationOssInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetQualificationOssInfoResponseBody) *GetQualificationOssInfoResponse
	GetBody() *GetQualificationOssInfoResponseBody
}

type GetQualificationOssInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualificationOssInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualificationOssInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationOssInfoResponse) GoString() string {
	return s.String()
}

func (s *GetQualificationOssInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualificationOssInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualificationOssInfoResponse) GetBody() *GetQualificationOssInfoResponseBody {
	return s.Body
}

func (s *GetQualificationOssInfoResponse) SetHeaders(v map[string]*string) *GetQualificationOssInfoResponse {
	s.Headers = v
	return s
}

func (s *GetQualificationOssInfoResponse) SetStatusCode(v int32) *GetQualificationOssInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualificationOssInfoResponse) SetBody(v *GetQualificationOssInfoResponseBody) *GetQualificationOssInfoResponse {
	s.Body = v
	return s
}

func (s *GetQualificationOssInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
