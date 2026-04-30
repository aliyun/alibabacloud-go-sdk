// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInstructionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableInstructionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableInstructionsResponse
	GetStatusCode() *int32
	SetBody(v *GetTableInstructionsResponseBody) *GetTableInstructionsResponse
	GetBody() *GetTableInstructionsResponseBody
}

type GetTableInstructionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableInstructionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableInstructionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableInstructionsResponse) GoString() string {
	return s.String()
}

func (s *GetTableInstructionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableInstructionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableInstructionsResponse) GetBody() *GetTableInstructionsResponseBody {
	return s.Body
}

func (s *GetTableInstructionsResponse) SetHeaders(v map[string]*string) *GetTableInstructionsResponse {
	s.Headers = v
	return s
}

func (s *GetTableInstructionsResponse) SetStatusCode(v int32) *GetTableInstructionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableInstructionsResponse) SetBody(v *GetTableInstructionsResponseBody) *GetTableInstructionsResponse {
	s.Body = v
	return s
}

func (s *GetTableInstructionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
