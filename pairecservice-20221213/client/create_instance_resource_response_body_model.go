// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInstanceResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *CreateInstanceResourceResponseBody
	GetResourceId() *string
}

type CreateInstanceResourceResponseBody struct {
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CreateInstanceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateInstanceResourceResponseBody) SetRequestId(v string) *CreateInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResourceResponseBody) SetResourceId(v string) *CreateInstanceResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateInstanceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
