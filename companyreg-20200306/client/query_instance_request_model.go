// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryInstanceRequest
	GetBizType() *string
	SetInstanceId(v string) *QueryInstanceRequest
	GetInstanceId() *string
}

type QueryInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// T20210302164856000001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceRequest) SetBizType(v string) *QueryInstanceRequest {
	s.BizType = &v
	return s
}

func (s *QueryInstanceRequest) SetInstanceId(v string) *QueryInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceRequest) Validate() error {
	return dara.Validate(s)
}
