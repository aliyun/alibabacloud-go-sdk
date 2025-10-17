// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodesParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodesParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodesParametersResponseBody) *ModifyDBNodesParametersResponse
	GetBody() *ModifyDBNodesParametersResponseBody
}

type ModifyDBNodesParametersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodesParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodesParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodesParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodesParametersResponse) GetBody() *ModifyDBNodesParametersResponseBody {
	return s.Body
}

func (s *ModifyDBNodesParametersResponse) SetHeaders(v map[string]*string) *ModifyDBNodesParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodesParametersResponse) SetStatusCode(v int32) *ModifyDBNodesParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodesParametersResponse) SetBody(v *ModifyDBNodesParametersResponseBody) *ModifyDBNodesParametersResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodesParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
