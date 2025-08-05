// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteControlPolicyTemplateResponseBody
	GetRequestId() *string
}

type DeleteControlPolicyTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4E7F94C7-781F-5192-86CF-DB085013C810
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteControlPolicyTemplateResponseBody) SetRequestId(v string) *DeleteControlPolicyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteControlPolicyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
