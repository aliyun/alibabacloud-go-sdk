// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileSummaryInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialFileSummaryInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialFileSummaryInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialFileSummaryInfoResponseBody) *QueryMaterialFileSummaryInfoResponse
	GetBody() *QueryMaterialFileSummaryInfoResponseBody
}

type QueryMaterialFileSummaryInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialFileSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialFileSummaryInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileSummaryInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialFileSummaryInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialFileSummaryInfoResponse) GetBody() *QueryMaterialFileSummaryInfoResponseBody {
	return s.Body
}

func (s *QueryMaterialFileSummaryInfoResponse) SetHeaders(v map[string]*string) *QueryMaterialFileSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponse) SetStatusCode(v int32) *QueryMaterialFileSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponse) SetBody(v *QueryMaterialFileSummaryInfoResponseBody) *QueryMaterialFileSummaryInfoResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
