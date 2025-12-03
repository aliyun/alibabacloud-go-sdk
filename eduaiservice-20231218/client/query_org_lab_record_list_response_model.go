// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgLabRecordListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrgLabRecordListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrgLabRecordListResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrgLabRecordListResponseBody) *QueryOrgLabRecordListResponse
	GetBody() *QueryOrgLabRecordListResponseBody
}

type QueryOrgLabRecordListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgLabRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgLabRecordListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgLabRecordListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgLabRecordListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrgLabRecordListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrgLabRecordListResponse) GetBody() *QueryOrgLabRecordListResponseBody {
	return s.Body
}

func (s *QueryOrgLabRecordListResponse) SetHeaders(v map[string]*string) *QueryOrgLabRecordListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgLabRecordListResponse) SetStatusCode(v int32) *QueryOrgLabRecordListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgLabRecordListResponse) SetBody(v *QueryOrgLabRecordListResponseBody) *QueryOrgLabRecordListResponse {
	s.Body = v
	return s
}

func (s *QueryOrgLabRecordListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
