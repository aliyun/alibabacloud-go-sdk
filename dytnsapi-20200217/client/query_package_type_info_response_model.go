// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPackageTypeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPackageTypeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPackageTypeInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryPackageTypeInfoResponseBody) *QueryPackageTypeInfoResponse
	GetBody() *QueryPackageTypeInfoResponseBody
}

type QueryPackageTypeInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPackageTypeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPackageTypeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPackageTypeInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPackageTypeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPackageTypeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPackageTypeInfoResponse) GetBody() *QueryPackageTypeInfoResponseBody {
	return s.Body
}

func (s *QueryPackageTypeInfoResponse) SetHeaders(v map[string]*string) *QueryPackageTypeInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPackageTypeInfoResponse) SetStatusCode(v int32) *QueryPackageTypeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPackageTypeInfoResponse) SetBody(v *QueryPackageTypeInfoResponseBody) *QueryPackageTypeInfoResponse {
	s.Body = v
	return s
}

func (s *QueryPackageTypeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
