// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStagingConfigToProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishStagingConfigToProductionResponseBody
	GetRequestId() *string
}

type PublishStagingConfigToProductionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishStagingConfigToProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishStagingConfigToProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
