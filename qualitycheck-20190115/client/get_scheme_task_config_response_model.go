// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSchemeTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSchemeTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSchemeTaskConfigResponseBody) *GetSchemeTaskConfigResponse
	GetBody() *GetSchemeTaskConfigResponseBody
}

type GetSchemeTaskConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSchemeTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSchemeTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSchemeTaskConfigResponse) GetBody() *GetSchemeTaskConfigResponseBody {
	return s.Body
}

func (s *GetSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *GetSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSchemeTaskConfigResponse) SetStatusCode(v int32) *GetSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchemeTaskConfigResponse) SetBody(v *GetSchemeTaskConfigResponseBody) *GetSchemeTaskConfigResponse {
	s.Body = v
	return s
}

func (s *GetSchemeTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
