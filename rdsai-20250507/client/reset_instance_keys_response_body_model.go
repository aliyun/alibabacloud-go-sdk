// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ResetInstanceKeysResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ResetInstanceKeysResponseBody
	GetRequestId() *string
}

type ResetInstanceKeysResponseBody struct {
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetInstanceKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstanceKeysResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ResetInstanceKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetInstanceKeysResponseBody) SetInstanceName(v string) *ResetInstanceKeysResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ResetInstanceKeysResponseBody) SetRequestId(v string) *ResetInstanceKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetInstanceKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
