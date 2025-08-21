// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAICInstanceImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeAICInstanceImageResponseBody
	GetRequestId() *string
}

type UpgradeAICInstanceImageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeAICInstanceImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAICInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAICInstanceImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAICInstanceImageResponseBody) SetRequestId(v string) *UpgradeAICInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAICInstanceImageResponseBody) Validate() error {
	return dara.Validate(s)
}
