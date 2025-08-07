// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseSceneConfigResponseBody
	GetRequestId() *string
}

type ModifyDefenseSceneConfigResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseSceneConfigResponseBody) SetRequestId(v string) *ModifyDefenseSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
