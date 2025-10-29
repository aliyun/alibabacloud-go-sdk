// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableCompactionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableCompactionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableCompactionResponse
	GetStatusCode() *int32
	SetBody(v *TableCompaction) *GetTableCompactionResponse
	GetBody() *TableCompaction
}

type GetTableCompactionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TableCompaction   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableCompactionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableCompactionResponse) GoString() string {
	return s.String()
}

func (s *GetTableCompactionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableCompactionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableCompactionResponse) GetBody() *TableCompaction {
	return s.Body
}

func (s *GetTableCompactionResponse) SetHeaders(v map[string]*string) *GetTableCompactionResponse {
	s.Headers = v
	return s
}

func (s *GetTableCompactionResponse) SetStatusCode(v int32) *GetTableCompactionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableCompactionResponse) SetBody(v *TableCompaction) *GetTableCompactionResponse {
	s.Body = v
	return s
}

func (s *GetTableCompactionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
