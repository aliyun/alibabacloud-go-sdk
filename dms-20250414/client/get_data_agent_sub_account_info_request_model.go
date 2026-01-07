// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentSubAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDmsUnit(v string) *GetDataAgentSubAccountInfoRequest
	GetDmsUnit() *string
	SetSubAccountId(v string) *GetDataAgentSubAccountInfoRequest
	GetSubAccountId() *string
}

type GetDataAgentSubAccountInfoRequest struct {
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// example:
	//
	// 20282*****7591
	SubAccountId *string `json:"SubAccountId,omitempty" xml:"SubAccountId,omitempty"`
}

func (s GetDataAgentSubAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentSubAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentSubAccountInfoRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *GetDataAgentSubAccountInfoRequest) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *GetDataAgentSubAccountInfoRequest) SetDmsUnit(v string) *GetDataAgentSubAccountInfoRequest {
	s.DmsUnit = &v
	return s
}

func (s *GetDataAgentSubAccountInfoRequest) SetSubAccountId(v string) *GetDataAgentSubAccountInfoRequest {
	s.SubAccountId = &v
	return s
}

func (s *GetDataAgentSubAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
