// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSqlInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneSqlInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneSqlInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CloneSqlInstanceResponseBody) *CloneSqlInstanceResponse
	GetBody() *CloneSqlInstanceResponseBody
}

type CloneSqlInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneSqlInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *CloneSqlInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneSqlInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneSqlInstanceResponse) GetBody() *CloneSqlInstanceResponseBody {
	return s.Body
}

func (s *CloneSqlInstanceResponse) SetHeaders(v map[string]*string) *CloneSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *CloneSqlInstanceResponse) SetStatusCode(v int32) *CloneSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneSqlInstanceResponse) SetBody(v *CloneSqlInstanceResponseBody) *CloneSqlInstanceResponse {
	s.Body = v
	return s
}

func (s *CloneSqlInstanceResponse) Validate() error {
	return dara.Validate(s)
}
