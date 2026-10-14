// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticTableIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticTableIndexResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticTableIndexResponseBody) *GetAgenticTableIndexResponse
	GetBody() *GetAgenticTableIndexResponseBody
}

type GetAgenticTableIndexResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticTableIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticTableIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableIndexResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticTableIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticTableIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticTableIndexResponse) GetBody() *GetAgenticTableIndexResponseBody {
	return s.Body
}

func (s *GetAgenticTableIndexResponse) SetHeaders(v map[string]*string) *GetAgenticTableIndexResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticTableIndexResponse) SetStatusCode(v int32) *GetAgenticTableIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticTableIndexResponse) SetBody(v *GetAgenticTableIndexResponseBody) *GetAgenticTableIndexResponse {
	s.Body = v
	return s
}

func (s *GetAgenticTableIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
