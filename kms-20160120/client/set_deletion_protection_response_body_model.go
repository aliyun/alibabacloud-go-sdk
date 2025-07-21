// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDeletionProtectionResponseBody
	GetRequestId() *string
}

type SetDeletionProtectionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDeletionProtectionResponseBody) SetRequestId(v string) *SetDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
