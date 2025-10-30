// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineageByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableColumnLineageByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableColumnLineageByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTableColumnLineageByTaskIdResponseBody) *GetTableColumnLineageByTaskIdResponse
	GetBody() *GetTableColumnLineageByTaskIdResponseBody
}

type GetTableColumnLineageByTaskIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableColumnLineageByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableColumnLineageByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableColumnLineageByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableColumnLineageByTaskIdResponse) GetBody() *GetTableColumnLineageByTaskIdResponseBody {
	return s.Body
}

func (s *GetTableColumnLineageByTaskIdResponse) SetHeaders(v map[string]*string) *GetTableColumnLineageByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponse) SetStatusCode(v int32) *GetTableColumnLineageByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponse) SetBody(v *GetTableColumnLineageByTaskIdResponseBody) *GetTableColumnLineageByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
