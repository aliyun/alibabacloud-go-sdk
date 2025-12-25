// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayNamespaceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPrepayNamespaceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPrepayNamespaceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPrepayNamespaceSpecResponseBody) *ModifyPrepayNamespaceSpecResponse
	GetBody() *ModifyPrepayNamespaceSpecResponseBody
}

type ModifyPrepayNamespaceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrepayNamespaceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrepayNamespaceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPrepayNamespaceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPrepayNamespaceSpecResponse) GetBody() *ModifyPrepayNamespaceSpecResponseBody {
	return s.Body
}

func (s *ModifyPrepayNamespaceSpecResponse) SetHeaders(v map[string]*string) *ModifyPrepayNamespaceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponse) SetStatusCode(v int32) *ModifyPrepayNamespaceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponse) SetBody(v *ModifyPrepayNamespaceSpecResponseBody) *ModifyPrepayNamespaceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
