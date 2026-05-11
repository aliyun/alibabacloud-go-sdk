// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTTSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTTSConfigResponseBody
	GetRequestId() *string
}

type ModifyTTSConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTTSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTTSConfigResponseBody) SetRequestId(v string) *ModifyTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
