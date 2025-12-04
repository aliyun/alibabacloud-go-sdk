// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSLBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSLBResponseBody
	GetRequestId() *string
}

type DeleteSLBResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSLBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSLBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSLBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSLBResponseBody) SetRequestId(v string) *DeleteSLBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSLBResponseBody) Validate() error {
	return dara.Validate(s)
}
