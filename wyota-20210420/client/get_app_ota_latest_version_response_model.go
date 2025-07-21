// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppOtaLatestVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppOtaLatestVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppOtaLatestVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetAppOtaLatestVersionResponseBody) *GetAppOtaLatestVersionResponse
	GetBody() *GetAppOtaLatestVersionResponseBody
}

type GetAppOtaLatestVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppOtaLatestVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppOtaLatestVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppOtaLatestVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppOtaLatestVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppOtaLatestVersionResponse) GetBody() *GetAppOtaLatestVersionResponseBody {
	return s.Body
}

func (s *GetAppOtaLatestVersionResponse) SetHeaders(v map[string]*string) *GetAppOtaLatestVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAppOtaLatestVersionResponse) SetStatusCode(v int32) *GetAppOtaLatestVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppOtaLatestVersionResponse) SetBody(v *GetAppOtaLatestVersionResponseBody) *GetAppOtaLatestVersionResponse {
	s.Body = v
	return s
}

func (s *GetAppOtaLatestVersionResponse) Validate() error {
	return dara.Validate(s)
}
