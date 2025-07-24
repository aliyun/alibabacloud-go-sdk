// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateScriptResponseBody
	GetRequestId() *string
}

type UpdateScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScriptResponseBody) SetRequestId(v string) *UpdateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
