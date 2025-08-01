// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteClusterRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *DeleteClusterRequest
	GetInstanceId() *string
}

type DeleteClusterRequest struct {
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
	// mse-cn-6ja1rgl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteClusterRequest) SetAcceptLanguage(v string) *DeleteClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteClusterRequest) SetInstanceId(v string) *DeleteClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	return dara.Validate(s)
}
