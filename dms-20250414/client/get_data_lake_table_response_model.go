// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeTableResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeTableResponseBody) *GetDataLakeTableResponse
	GetBody() *GetDataLakeTableResponseBody
}

type GetDataLakeTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeTableResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeTableResponse) GetBody() *GetDataLakeTableResponseBody {
	return s.Body
}

func (s *GetDataLakeTableResponse) SetHeaders(v map[string]*string) *GetDataLakeTableResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeTableResponse) SetStatusCode(v int32) *GetDataLakeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeTableResponse) SetBody(v *GetDataLakeTableResponseBody) *GetDataLakeTableResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
