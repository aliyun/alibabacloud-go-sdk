// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVodStagingConfigToProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishVodStagingConfigToProductionResponseBody
	GetRequestId() *string
}

type PublishVodStagingConfigToProductionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishVodStagingConfigToProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishVodStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishVodStagingConfigToProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishVodStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishVodStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishVodStagingConfigToProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
