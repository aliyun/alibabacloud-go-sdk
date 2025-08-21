// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElastictaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *ModifyElastictaskRequest
	GetBody() *string
}

type ModifyElastictaskRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElastictaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskRequest) GetBody() *string {
	return s.Body
}

func (s *ModifyElastictaskRequest) SetBody(v string) *ModifyElastictaskRequest {
	s.Body = &v
	return s
}

func (s *ModifyElastictaskRequest) Validate() error {
	return dara.Validate(s)
}
