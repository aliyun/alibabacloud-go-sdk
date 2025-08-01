// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RetryClusterRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *RetryClusterRequest
	GetInstanceId() *string
	SetRequestPars(v string) *RetryClusterRequest
	GetRequestPars() *string
}

type RetryClusterRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s RetryClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryClusterRequest) GoString() string {
	return s.String()
}

func (s *RetryClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RetryClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RetryClusterRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *RetryClusterRequest) SetAcceptLanguage(v string) *RetryClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RetryClusterRequest) SetInstanceId(v string) *RetryClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *RetryClusterRequest) SetRequestPars(v string) *RetryClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *RetryClusterRequest) Validate() error {
	return dara.Validate(s)
}
