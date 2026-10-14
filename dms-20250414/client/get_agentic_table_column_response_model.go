// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableColumnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticTableColumnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticTableColumnResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticTableColumnResponseBody) *GetAgenticTableColumnResponse
	GetBody() *GetAgenticTableColumnResponseBody
}

type GetAgenticTableColumnResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticTableColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticTableColumnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableColumnResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticTableColumnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticTableColumnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticTableColumnResponse) GetBody() *GetAgenticTableColumnResponseBody {
	return s.Body
}

func (s *GetAgenticTableColumnResponse) SetHeaders(v map[string]*string) *GetAgenticTableColumnResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticTableColumnResponse) SetStatusCode(v int32) *GetAgenticTableColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticTableColumnResponse) SetBody(v *GetAgenticTableColumnResponseBody) *GetAgenticTableColumnResponse {
	s.Body = v
	return s
}

func (s *GetAgenticTableColumnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
