// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoTagScanStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoTagScanStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoTagScanStatusResponseBody) *GetRepoTagScanStatusResponse
	GetBody() *GetRepoTagScanStatusResponseBody
}

type GetRepoTagScanStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagScanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagScanStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoTagScanStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoTagScanStatusResponse) GetBody() *GetRepoTagScanStatusResponseBody {
	return s.Body
}

func (s *GetRepoTagScanStatusResponse) SetHeaders(v map[string]*string) *GetRepoTagScanStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagScanStatusResponse) SetStatusCode(v int32) *GetRepoTagScanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagScanStatusResponse) SetBody(v *GetRepoTagScanStatusResponseBody) *GetRepoTagScanStatusResponse {
	s.Body = v
	return s
}

func (s *GetRepoTagScanStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
