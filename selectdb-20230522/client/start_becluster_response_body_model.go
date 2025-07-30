// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBEClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartBEClusterResponseBody
	GetRequestId() *string
}

type StartBEClusterResponseBody struct {
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBEClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartBEClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBEClusterResponseBody) SetRequestId(v string) *StartBEClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBEClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
