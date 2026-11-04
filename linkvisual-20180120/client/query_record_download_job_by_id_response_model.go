// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadJobByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordDownloadJobByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordDownloadJobByIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordDownloadJobByIdResponseBody) *QueryRecordDownloadJobByIdResponse
	GetBody() *QueryRecordDownloadJobByIdResponseBody
}

type QueryRecordDownloadJobByIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadJobByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadJobByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadJobByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordDownloadJobByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordDownloadJobByIdResponse) GetBody() *QueryRecordDownloadJobByIdResponseBody {
	return s.Body
}

func (s *QueryRecordDownloadJobByIdResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadJobByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadJobByIdResponse) SetStatusCode(v int32) *QueryRecordDownloadJobByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadJobByIdResponse) SetBody(v *QueryRecordDownloadJobByIdResponseBody) *QueryRecordDownloadJobByIdResponse {
	s.Body = v
	return s
}

func (s *QueryRecordDownloadJobByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
