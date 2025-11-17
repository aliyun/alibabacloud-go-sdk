// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetDetailInfoResponseBody) *QueryDatasetDetailInfoResponse
	GetBody() *QueryDatasetDetailInfoResponseBody
}

type QueryDatasetDetailInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetDetailInfoResponse) GetBody() *QueryDatasetDetailInfoResponseBody {
	return s.Body
}

func (s *QueryDatasetDetailInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetDetailInfoResponse) SetStatusCode(v int32) *QueryDatasetDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetDetailInfoResponse) SetBody(v *QueryDatasetDetailInfoResponseBody) *QueryDatasetDetailInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetDetailInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
