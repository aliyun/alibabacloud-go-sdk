// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableColumnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableColumnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableColumnResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableColumnResponseBody) *GetMetaTableColumnResponse
	GetBody() *GetMetaTableColumnResponseBody
}

type GetMetaTableColumnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableColumnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableColumnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableColumnResponse) GetBody() *GetMetaTableColumnResponseBody {
	return s.Body
}

func (s *GetMetaTableColumnResponse) SetHeaders(v map[string]*string) *GetMetaTableColumnResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableColumnResponse) SetStatusCode(v int32) *GetMetaTableColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableColumnResponse) SetBody(v *GetMetaTableColumnResponseBody) *GetMetaTableColumnResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableColumnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
