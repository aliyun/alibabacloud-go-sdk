// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordHistoryConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPasswordHistoryConfigurationRequest
	GetInstanceId() *string
}

type GetPasswordHistoryConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordHistoryConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordHistoryConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPasswordHistoryConfigurationRequest) SetInstanceId(v string) *GetPasswordHistoryConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPasswordHistoryConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
