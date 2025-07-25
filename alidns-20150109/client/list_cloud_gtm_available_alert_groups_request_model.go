// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAvailableAlertGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmAvailableAlertGroupsRequest
	GetAcceptLanguage() *string
}

type ListCloudGtmAvailableAlertGroupsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US**: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListCloudGtmAvailableAlertGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAvailableAlertGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAvailableAlertGroupsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmAvailableAlertGroupsRequest) SetAcceptLanguage(v string) *ListCloudGtmAvailableAlertGroupsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsRequest) Validate() error {
	return dara.Validate(s)
}
