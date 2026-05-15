// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHotlineServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *StartHotlineServiceRequest
	GetAccountName() *string
	SetClientToken(v string) *StartHotlineServiceRequest
	GetClientToken() *string
	SetInstanceId(v string) *StartHotlineServiceRequest
	GetInstanceId() *string
}

type StartHotlineServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartHotlineServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *StartHotlineServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartHotlineServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartHotlineServiceRequest) SetAccountName(v string) *StartHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *StartHotlineServiceRequest) SetClientToken(v string) *StartHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartHotlineServiceRequest) SetInstanceId(v string) *StartHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartHotlineServiceRequest) Validate() error {
	return dara.Validate(s)
}
