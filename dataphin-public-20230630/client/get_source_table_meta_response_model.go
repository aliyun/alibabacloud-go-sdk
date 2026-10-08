// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourceTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourceTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetSourceTableMetaResponseBody) *GetSourceTableMetaResponse
	GetBody() *GetSourceTableMetaResponseBody
}

type GetSourceTableMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourceTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourceTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaResponse) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourceTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourceTableMetaResponse) GetBody() *GetSourceTableMetaResponseBody {
	return s.Body
}

func (s *GetSourceTableMetaResponse) SetHeaders(v map[string]*string) *GetSourceTableMetaResponse {
	s.Headers = v
	return s
}

func (s *GetSourceTableMetaResponse) SetStatusCode(v int32) *GetSourceTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceTableMetaResponse) SetBody(v *GetSourceTableMetaResponseBody) *GetSourceTableMetaResponse {
	s.Body = v
	return s
}

func (s *GetSourceTableMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
