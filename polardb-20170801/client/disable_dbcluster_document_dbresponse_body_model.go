// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDocumentDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDBClusterDocumentDBResponseBody
	GetRequestId() *string
}

type DisableDBClusterDocumentDBResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBClusterDocumentDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDocumentDBResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDocumentDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBClusterDocumentDBResponseBody) SetRequestId(v string) *DisableDBClusterDocumentDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBClusterDocumentDBResponseBody) Validate() error {
	return dara.Validate(s)
}
