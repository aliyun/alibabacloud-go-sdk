// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectClientRuleResponseBody
	GetRequestId() *string
}

type UpdateFileProtectClientRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectClientRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectClientRuleResponseBody) SetRequestId(v string) *UpdateFileProtectClientRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectClientRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
