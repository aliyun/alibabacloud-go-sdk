// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GrantSagInstanceToVbrResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GrantSagInstanceToVbrResponseBody
	GetRequestId() *string
}

type GrantSagInstanceToVbrResponseBody struct {
	// The ID of the authorization.
	//
	// example:
	//
	// sgv-3x8djyem7vqh70****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 49CEBB2B-9E5C-4789-8A29-3255A56A67B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantSagInstanceToVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToVbrResponseBody) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToVbrResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantSagInstanceToVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantSagInstanceToVbrResponseBody) SetInstanceId(v string) *GrantSagInstanceToVbrResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GrantSagInstanceToVbrResponseBody) SetRequestId(v string) *GrantSagInstanceToVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantSagInstanceToVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
