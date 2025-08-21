// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteLogstashRequest
	GetClientToken() *string
	SetDeleteType(v string) *DeleteLogstashRequest
	GetDeleteType() *string
}

type DeleteLogstashRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The type of the release operation. Valid values:
	//
	// 	- immediate: The cluster is immediately deleted when it is released. After the cluster is deleted, the data stored in the cluster is deleted, and the system removes the cluster from the Logstash cluster list.
	//
	// 	- protective: The cluster is released 24 hours later. During the period of 24 hours, you can still find the cluster in the Logstash cluster list, and [restore the cluster](https://help.aliyun.com/document_detail/202205.html) or [immediately release the cluster](https://help.aliyun.com/document_detail/160591.html). After 24 hours elapse, the data stored in the cluster is deleted.
	//
	// example:
	//
	// protective
	DeleteType *string `json:"deleteType,omitempty" xml:"deleteType,omitempty"`
}

func (s DeleteLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogstashRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogstashRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteLogstashRequest) GetDeleteType() *string {
	return s.DeleteType
}

func (s *DeleteLogstashRequest) SetClientToken(v string) *DeleteLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLogstashRequest) SetDeleteType(v string) *DeleteLogstashRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteLogstashRequest) Validate() error {
	return dara.Validate(s)
}
