// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectClientRuleStatusResponseBody
	GetRequestId() *string
}

type UpdateFileProtectClientRuleStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectClientRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectClientRuleStatusResponseBody) SetRequestId(v string) *UpdateFileProtectClientRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
