// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectClientRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFileProtectClientRuleResponseBody
	GetRequestId() *string
}

type DeleteFileProtectClientRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileProtectClientRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectClientRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectClientRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileProtectClientRuleResponseBody) SetRequestId(v string) *DeleteFileProtectClientRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileProtectClientRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
