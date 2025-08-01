// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutClusterHealthCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *PutClusterHealthCheckTaskRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *PutClusterHealthCheckTaskRequest
	GetInstanceId() *string
}

type PutClusterHealthCheckTaskRequest struct {
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
	// mse_prepaid_public_cn-2r42o83h506
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PutClusterHealthCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PutClusterHealthCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *PutClusterHealthCheckTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *PutClusterHealthCheckTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PutClusterHealthCheckTaskRequest) SetAcceptLanguage(v string) *PutClusterHealthCheckTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *PutClusterHealthCheckTaskRequest) SetInstanceId(v string) *PutClusterHealthCheckTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *PutClusterHealthCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
