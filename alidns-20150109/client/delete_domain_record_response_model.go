// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainRecordResponseBody) *DeleteDomainRecordResponse
	GetBody() *DeleteDomainRecordResponseBody
}

type DeleteDomainRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainRecordResponse) GetBody() *DeleteDomainRecordResponseBody {
	return s.Body
}

func (s *DeleteDomainRecordResponse) SetHeaders(v map[string]*string) *DeleteDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainRecordResponse) SetStatusCode(v int32) *DeleteDomainRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainRecordResponse) SetBody(v *DeleteDomainRecordResponseBody) *DeleteDomainRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainRecordResponse) Validate() error {
	return dara.Validate(s)
}
