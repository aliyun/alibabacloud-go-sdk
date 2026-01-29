// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDPUtilizationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDPUtilizationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDPUtilizationDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryDPUtilizationDetailResponseBody) *QueryDPUtilizationDetailResponse
	GetBody() *QueryDPUtilizationDetailResponseBody
}

type QueryDPUtilizationDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDPUtilizationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDPUtilizationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDPUtilizationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDPUtilizationDetailResponse) GetBody() *QueryDPUtilizationDetailResponseBody {
	return s.Body
}

func (s *QueryDPUtilizationDetailResponse) SetHeaders(v map[string]*string) *QueryDPUtilizationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryDPUtilizationDetailResponse) SetStatusCode(v int32) *QueryDPUtilizationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDPUtilizationDetailResponse) SetBody(v *QueryDPUtilizationDetailResponseBody) *QueryDPUtilizationDetailResponse {
	s.Body = v
	return s
}

func (s *QueryDPUtilizationDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
