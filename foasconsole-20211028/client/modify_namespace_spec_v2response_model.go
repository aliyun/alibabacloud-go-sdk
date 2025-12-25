// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNamespaceSpecV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNamespaceSpecV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNamespaceSpecV2Response
	GetStatusCode() *int32
	SetBody(v *ModifyNamespaceSpecV2ResponseBody) *ModifyNamespaceSpecV2Response
	GetBody() *ModifyNamespaceSpecV2ResponseBody
}

type ModifyNamespaceSpecV2Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNamespaceSpecV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNamespaceSpecV2Response) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2Response) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNamespaceSpecV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNamespaceSpecV2Response) GetBody() *ModifyNamespaceSpecV2ResponseBody {
	return s.Body
}

func (s *ModifyNamespaceSpecV2Response) SetHeaders(v map[string]*string) *ModifyNamespaceSpecV2Response {
	s.Headers = v
	return s
}

func (s *ModifyNamespaceSpecV2Response) SetStatusCode(v int32) *ModifyNamespaceSpecV2Response {
	s.StatusCode = &v
	return s
}

func (s *ModifyNamespaceSpecV2Response) SetBody(v *ModifyNamespaceSpecV2ResponseBody) *ModifyNamespaceSpecV2Response {
	s.Body = v
	return s
}

func (s *ModifyNamespaceSpecV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
