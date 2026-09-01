// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckScopeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckScopeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckScopeConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckScopeConfigResponseBody) *GetCheckScopeConfigResponse
	GetBody() *GetCheckScopeConfigResponseBody
}

type GetCheckScopeConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckScopeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckScopeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckScopeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetCheckScopeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckScopeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckScopeConfigResponse) GetBody() *GetCheckScopeConfigResponseBody {
	return s.Body
}

func (s *GetCheckScopeConfigResponse) SetHeaders(v map[string]*string) *GetCheckScopeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetCheckScopeConfigResponse) SetStatusCode(v int32) *GetCheckScopeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckScopeConfigResponse) SetBody(v *GetCheckScopeConfigResponseBody) *GetCheckScopeConfigResponse {
	s.Body = v
	return s
}

func (s *GetCheckScopeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
