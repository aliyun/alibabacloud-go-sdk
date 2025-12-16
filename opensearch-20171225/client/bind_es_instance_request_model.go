// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *BindEsInstanceRequest
	GetBody() map[string]interface{}
}

type BindEsInstanceRequest struct {
	// The body of the request.
	//
	// example:
	//
	// {
	//
	//   "esInstanceId": "es-cn-abcde"
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindEsInstanceRequest) GoString() string {
	return s.String()
}

func (s *BindEsInstanceRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *BindEsInstanceRequest) SetBody(v map[string]interface{}) *BindEsInstanceRequest {
	s.Body = v
	return s
}

func (s *BindEsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
