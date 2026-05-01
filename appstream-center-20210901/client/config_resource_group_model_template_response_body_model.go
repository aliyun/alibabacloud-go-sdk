// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigResourceGroupModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigResourceGroupModelTemplateResponseBody
	GetRequestId() *string
}

type ConfigResourceGroupModelTemplateResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigResourceGroupModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigResourceGroupModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigResourceGroupModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigResourceGroupModelTemplateResponseBody) SetRequestId(v string) *ConfigResourceGroupModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
