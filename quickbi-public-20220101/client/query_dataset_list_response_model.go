// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetListResponseBody) *QueryDatasetListResponse
	GetBody() *QueryDatasetListResponseBody
}

type QueryDatasetListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetListResponse) GetBody() *QueryDatasetListResponseBody {
	return s.Body
}

func (s *QueryDatasetListResponse) SetHeaders(v map[string]*string) *QueryDatasetListResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetListResponse) SetStatusCode(v int32) *QueryDatasetListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetListResponse) SetBody(v *QueryDatasetListResponseBody) *QueryDatasetListResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
