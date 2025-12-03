// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNetTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterNetTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterNetTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterNetTypeResponseBody) *ModifyClusterNetTypeResponse
	GetBody() *ModifyClusterNetTypeResponseBody
}

type ModifyClusterNetTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterNetTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNetTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterNetTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterNetTypeResponse) GetBody() *ModifyClusterNetTypeResponseBody {
	return s.Body
}

func (s *ModifyClusterNetTypeResponse) SetHeaders(v map[string]*string) *ModifyClusterNetTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNetTypeResponse) SetStatusCode(v int32) *ModifyClusterNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterNetTypeResponse) SetBody(v *ModifyClusterNetTypeResponseBody) *ModifyClusterNetTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterNetTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
