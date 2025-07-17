// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetGroupDeletionProtectionResponseBody
	GetRequestId() *string
}

type SetGroupDeletionProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CCC29E24-3AEC-4F2C-8A14-78B14FA738B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGroupDeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetGroupDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetGroupDeletionProtectionResponseBody) SetRequestId(v string) *SetGroupDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGroupDeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
