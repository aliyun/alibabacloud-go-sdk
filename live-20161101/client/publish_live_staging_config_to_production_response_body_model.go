// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishLiveStagingConfigToProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishLiveStagingConfigToProductionResponseBody
	GetRequestId() *string
}

type PublishLiveStagingConfigToProductionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishLiveStagingConfigToProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishLiveStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishLiveStagingConfigToProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishLiveStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishLiveStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
