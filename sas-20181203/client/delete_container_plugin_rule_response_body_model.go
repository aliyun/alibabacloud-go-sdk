// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerPluginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteContainerPluginRuleResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteContainerPluginRuleResponseBody
	GetRequestId() *string
}

type DeleteContainerPluginRuleResponseBody struct {
	// Indicates whether the container escape prevention rule is deleted. Valid values:
	//
	// - **true**: The rule is deleted.
	//
	// - **false**: The rule failed to be deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerPluginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerPluginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerPluginRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteContainerPluginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContainerPluginRuleResponseBody) SetData(v bool) *DeleteContainerPluginRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteContainerPluginRuleResponseBody) SetRequestId(v string) *DeleteContainerPluginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContainerPluginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
