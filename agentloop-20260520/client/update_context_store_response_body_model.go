// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateContextStoreResponseBody
	GetRequestId() *string
	SetStrategyVersion(v int32) *UpdateContextStoreResponseBody
	GetStrategyVersion() *int32
}

type UpdateContextStoreResponseBody struct {
	// The request ID, which is used to locate the request during troubleshooting.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The effective strategy version number after the update for the memory type. If the strategy remains unchanged, the version number is the same as before the update.
	//
	// example:
	//
	// 2
	StrategyVersion *int32 `json:"strategyVersion,omitempty" xml:"strategyVersion,omitempty"`
}

func (s UpdateContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextStoreResponseBody) GetStrategyVersion() *int32 {
	return s.StrategyVersion
}

func (s *UpdateContextStoreResponseBody) SetRequestId(v string) *UpdateContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextStoreResponseBody) SetStrategyVersion(v int32) *UpdateContextStoreResponseBody {
	s.StrategyVersion = &v
	return s
}

func (s *UpdateContextStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
