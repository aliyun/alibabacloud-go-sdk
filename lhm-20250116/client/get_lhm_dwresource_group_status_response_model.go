// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmDWResourceGroupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLhmDWResourceGroupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLhmDWResourceGroupStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetLhmDWResourceGroupStatusResponseBody) *GetLhmDWResourceGroupStatusResponse
	GetBody() *GetLhmDWResourceGroupStatusResponseBody
}

type GetLhmDWResourceGroupStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLhmDWResourceGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLhmDWResourceGroupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLhmDWResourceGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *GetLhmDWResourceGroupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLhmDWResourceGroupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLhmDWResourceGroupStatusResponse) GetBody() *GetLhmDWResourceGroupStatusResponseBody {
	return s.Body
}

func (s *GetLhmDWResourceGroupStatusResponse) SetHeaders(v map[string]*string) *GetLhmDWResourceGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponse) SetStatusCode(v int32) *GetLhmDWResourceGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponse) SetBody(v *GetLhmDWResourceGroupStatusResponseBody) *GetLhmDWResourceGroupStatusResponse {
	s.Body = v
	return s
}

func (s *GetLhmDWResourceGroupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
