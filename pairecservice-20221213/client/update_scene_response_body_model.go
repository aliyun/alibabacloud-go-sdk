// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSceneResponseBody
	GetRequestId() *string
}

type UpdateSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FC17887E-3C82-5096-8AA6-F4C2E7417245
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSceneResponseBody) SetRequestId(v string) *UpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
