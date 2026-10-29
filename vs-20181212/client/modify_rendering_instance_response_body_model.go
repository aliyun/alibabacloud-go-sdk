// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRenderingInstanceResponseBody
	GetRequestId() *string
}

type ModifyRenderingInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6DFE7B89-8532-566F-B5CE-924B10FCE7AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRenderingInstanceResponseBody) SetRequestId(v string) *ModifyRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
