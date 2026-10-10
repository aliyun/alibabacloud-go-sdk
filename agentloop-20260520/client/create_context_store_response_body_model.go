// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateContextStoreResponseBody
	GetRequestId() *string
	SetStrategyVersion(v int32) *CreateContextStoreResponseBody
	GetStrategyVersion() *int32
}

type CreateContextStoreResponseBody struct {
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	StrategyVersion *int32 `json:"strategyVersion,omitempty" xml:"strategyVersion,omitempty"`
}

func (s CreateContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextStoreResponseBody) GetStrategyVersion() *int32 {
	return s.StrategyVersion
}

func (s *CreateContextStoreResponseBody) SetRequestId(v string) *CreateContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextStoreResponseBody) SetStrategyVersion(v int32) *CreateContextStoreResponseBody {
	s.StrategyVersion = &v
	return s
}

func (s *CreateContextStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
