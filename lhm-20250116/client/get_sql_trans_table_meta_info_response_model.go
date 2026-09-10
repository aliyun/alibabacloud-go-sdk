// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTransTableMetaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlTransTableMetaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlTransTableMetaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlTransTableMetaInfoResponseBody) *GetSqlTransTableMetaInfoResponse
	GetBody() *GetSqlTransTableMetaInfoResponseBody
}

type GetSqlTransTableMetaInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlTransTableMetaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlTransTableMetaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTransTableMetaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSqlTransTableMetaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlTransTableMetaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlTransTableMetaInfoResponse) GetBody() *GetSqlTransTableMetaInfoResponseBody {
	return s.Body
}

func (s *GetSqlTransTableMetaInfoResponse) SetHeaders(v map[string]*string) *GetSqlTransTableMetaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSqlTransTableMetaInfoResponse) SetStatusCode(v int32) *GetSqlTransTableMetaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlTransTableMetaInfoResponse) SetBody(v *GetSqlTransTableMetaInfoResponseBody) *GetSqlTransTableMetaInfoResponse {
	s.Body = v
	return s
}

func (s *GetSqlTransTableMetaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
