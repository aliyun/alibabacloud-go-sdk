// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhatAppTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListWhatAppTemplateRequest
	GetCustSpaceId() *string
	SetHetuParams(v string) *ListWhatAppTemplateRequest
	GetHetuParams() *string
}

type ListWhatAppTemplateRequest struct {
	// example:
	//
	// 2983883892
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// {}
	HetuParams *string `json:"HetuParams,omitempty" xml:"HetuParams,omitempty"`
}

func (s ListWhatAppTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWhatAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListWhatAppTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListWhatAppTemplateRequest) GetHetuParams() *string {
	return s.HetuParams
}

func (s *ListWhatAppTemplateRequest) SetCustSpaceId(v string) *ListWhatAppTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListWhatAppTemplateRequest) SetHetuParams(v string) *ListWhatAppTemplateRequest {
	s.HetuParams = &v
	return s
}

func (s *ListWhatAppTemplateRequest) Validate() error {
	return dara.Validate(s)
}
