// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetTagStatus(v bool) *CreateInstanceResponseBody
	GetTagStatus() *bool
}

type CreateInstanceResponseBody struct {
	// Instance ID
	//
	// example:
	//
	// apigateway-hz-b3c5dadd5***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Request ID
	//
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the tag exists. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// True
	TagStatus *bool `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetTagStatus() *bool {
	return s.TagStatus
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetTagStatus(v bool) *CreateInstanceResponseBody {
	s.TagStatus = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
