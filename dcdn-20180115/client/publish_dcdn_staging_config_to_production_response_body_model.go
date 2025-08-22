// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDcdnStagingConfigToProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishDcdnStagingConfigToProductionResponseBody
	GetRequestId() *string
}

type PublishDcdnStagingConfigToProductionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDcdnStagingConfigToProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishDcdnStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishDcdnStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
