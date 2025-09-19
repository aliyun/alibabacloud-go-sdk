// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskRankInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDockerhubImageRiskRankInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDockerhubImageRiskRankInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDockerhubImageRiskRankInfoResponseBody) *GetDockerhubImageRiskRankInfoResponse
	GetBody() *GetDockerhubImageRiskRankInfoResponseBody
}

type GetDockerhubImageRiskRankInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDockerhubImageRiskRankInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDockerhubImageRiskRankInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDockerhubImageRiskRankInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDockerhubImageRiskRankInfoResponse) GetBody() *GetDockerhubImageRiskRankInfoResponseBody {
	return s.Body
}

func (s *GetDockerhubImageRiskRankInfoResponse) SetHeaders(v map[string]*string) *GetDockerhubImageRiskRankInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponse) SetStatusCode(v int32) *GetDockerhubImageRiskRankInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponse) SetBody(v *GetDockerhubImageRiskRankInfoResponseBody) *GetDockerhubImageRiskRankInfoResponse {
	s.Body = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponse) Validate() error {
	return dara.Validate(s)
}
