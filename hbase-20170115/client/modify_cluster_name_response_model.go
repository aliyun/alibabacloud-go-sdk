// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterNameResponseBody) *ModifyClusterNameResponse
	GetBody() *ModifyClusterNameResponseBody
}

type ModifyClusterNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterNameResponse) GetBody() *ModifyClusterNameResponseBody {
	return s.Body
}

func (s *ModifyClusterNameResponse) SetHeaders(v map[string]*string) *ModifyClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNameResponse) SetStatusCode(v int32) *ModifyClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterNameResponse) SetBody(v *ModifyClusterNameResponseBody) *ModifyClusterNameResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
