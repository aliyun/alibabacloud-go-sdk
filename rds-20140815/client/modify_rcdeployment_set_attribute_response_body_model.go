// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDeploymentSetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCDeploymentSetAttributeResponseBody
	GetRequestId() *string
}

type ModifyRCDeploymentSetAttributeResponseBody struct {
	// example:
	//
	// B61C08E5-403A-46A2-96C1-F7B1216DB10C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCDeploymentSetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDeploymentSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCDeploymentSetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCDeploymentSetAttributeResponseBody) SetRequestId(v string) *ModifyRCDeploymentSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
