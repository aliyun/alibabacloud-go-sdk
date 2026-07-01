// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticDBBranchResponseBody
	GetRequestId() *string
}

type DeleteAgenticDBBranchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D4E5F6A7-B8C9-0123-DEFA-456789012DEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticDBBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDBBranchResponseBody) SetRequestId(v string) *DeleteAgenticDBBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDBBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
