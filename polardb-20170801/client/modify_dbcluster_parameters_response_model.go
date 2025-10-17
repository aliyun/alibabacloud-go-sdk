// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterParametersResponseBody) *ModifyDBClusterParametersResponse
	GetBody() *ModifyDBClusterParametersResponseBody
}

type ModifyDBClusterParametersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterParametersResponse) GetBody() *ModifyDBClusterParametersResponseBody {
	return s.Body
}

func (s *ModifyDBClusterParametersResponse) SetHeaders(v map[string]*string) *ModifyDBClusterParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterParametersResponse) SetStatusCode(v int32) *ModifyDBClusterParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterParametersResponse) SetBody(v *ModifyDBClusterParametersResponseBody) *ModifyDBClusterParametersResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
