// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbakTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlFlashbakTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlFlashbakTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlFlashbakTaskResponseBody) *DescribeSqlFlashbakTaskResponse
	GetBody() *DescribeSqlFlashbakTaskResponseBody
}

type DescribeSqlFlashbakTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlFlashbakTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlFlashbakTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlFlashbakTaskResponse) GetBody() *DescribeSqlFlashbakTaskResponseBody {
	return s.Body
}

func (s *DescribeSqlFlashbakTaskResponse) SetHeaders(v map[string]*string) *DescribeSqlFlashbakTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetStatusCode(v int32) *DescribeSqlFlashbakTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) SetBody(v *DescribeSqlFlashbakTaskResponseBody) *DescribeSqlFlashbakTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponse) Validate() error {
	return dara.Validate(s)
}
