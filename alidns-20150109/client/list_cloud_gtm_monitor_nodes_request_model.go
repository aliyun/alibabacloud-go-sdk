// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmMonitorNodesRequest
	GetAcceptLanguage() *string
}

type ListCloudGtmMonitorNodesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListCloudGtmMonitorNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmMonitorNodesRequest) SetAcceptLanguage(v string) *ListCloudGtmMonitorNodesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmMonitorNodesRequest) Validate() error {
	return dara.Validate(s)
}
