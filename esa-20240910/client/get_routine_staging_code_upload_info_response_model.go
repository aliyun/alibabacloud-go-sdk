// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineStagingCodeUploadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineStagingCodeUploadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineStagingCodeUploadInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineStagingCodeUploadInfoResponseBody) *GetRoutineStagingCodeUploadInfoResponse
	GetBody() *GetRoutineStagingCodeUploadInfoResponseBody
}

type GetRoutineStagingCodeUploadInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineStagingCodeUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineStagingCodeUploadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineStagingCodeUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineStagingCodeUploadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineStagingCodeUploadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineStagingCodeUploadInfoResponse) GetBody() *GetRoutineStagingCodeUploadInfoResponseBody {
	return s.Body
}

func (s *GetRoutineStagingCodeUploadInfoResponse) SetHeaders(v map[string]*string) *GetRoutineStagingCodeUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponse) SetStatusCode(v int32) *GetRoutineStagingCodeUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponse) SetBody(v *GetRoutineStagingCodeUploadInfoResponseBody) *GetRoutineStagingCodeUploadInfoResponse {
	s.Body = v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
