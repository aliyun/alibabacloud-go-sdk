// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainDkimRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDomainDkimRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDomainDkimRecordResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDomainDkimRecordResponseBody) *ChangeDomainDkimRecordResponse
	GetBody() *ChangeDomainDkimRecordResponseBody
}

type ChangeDomainDkimRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDomainDkimRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDomainDkimRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainDkimRecordResponse) GoString() string {
	return s.String()
}

func (s *ChangeDomainDkimRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDomainDkimRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDomainDkimRecordResponse) GetBody() *ChangeDomainDkimRecordResponseBody {
	return s.Body
}

func (s *ChangeDomainDkimRecordResponse) SetHeaders(v map[string]*string) *ChangeDomainDkimRecordResponse {
	s.Headers = v
	return s
}

func (s *ChangeDomainDkimRecordResponse) SetStatusCode(v int32) *ChangeDomainDkimRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDomainDkimRecordResponse) SetBody(v *ChangeDomainDkimRecordResponseBody) *ChangeDomainDkimRecordResponse {
	s.Body = v
	return s
}

func (s *ChangeDomainDkimRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
