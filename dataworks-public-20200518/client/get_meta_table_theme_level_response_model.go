// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableThemeLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableThemeLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableThemeLevelResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableThemeLevelResponseBody) *GetMetaTableThemeLevelResponse
	GetBody() *GetMetaTableThemeLevelResponseBody
}

type GetMetaTableThemeLevelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableThemeLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableThemeLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableThemeLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableThemeLevelResponse) GetBody() *GetMetaTableThemeLevelResponseBody {
	return s.Body
}

func (s *GetMetaTableThemeLevelResponse) SetHeaders(v map[string]*string) *GetMetaTableThemeLevelResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableThemeLevelResponse) SetStatusCode(v int32) *GetMetaTableThemeLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableThemeLevelResponse) SetBody(v *GetMetaTableThemeLevelResponseBody) *GetMetaTableThemeLevelResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableThemeLevelResponse) Validate() error {
	return dara.Validate(s)
}
