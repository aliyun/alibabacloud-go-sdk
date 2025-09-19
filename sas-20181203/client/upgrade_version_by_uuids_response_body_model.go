// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionByUuidsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeVersionByUuidsResponseBody
	GetRequestId() *string
}

type UpgradeVersionByUuidsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeVersionByUuidsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionByUuidsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeVersionByUuidsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeVersionByUuidsResponseBody) SetRequestId(v string) *UpgradeVersionByUuidsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeVersionByUuidsResponseBody) Validate() error {
	return dara.Validate(s)
}
