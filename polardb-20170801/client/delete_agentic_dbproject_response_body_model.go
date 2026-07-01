// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticDBProjectResponseBody
	GetRequestId() *string
}

type DeleteAgenticDBProjectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6A7B8C9-D0E1-2345-FABC-456789012345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticDBProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDBProjectResponseBody) SetRequestId(v string) *DeleteAgenticDBProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDBProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
