// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrayDomainFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v []*GetGrayDomainFunctionResponseBodyDomainConfigList) *GetGrayDomainFunctionResponseBody
	GetDomainConfigList() []*GetGrayDomainFunctionResponseBodyDomainConfigList
	SetDomainName(v string) *GetGrayDomainFunctionResponseBody
	GetDomainName() *string
	SetProgression(v string) *GetGrayDomainFunctionResponseBody
	GetProgression() *string
	SetRequestId(v string) *GetGrayDomainFunctionResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetGrayDomainFunctionResponseBody
	GetStatus() *string
}

type GetGrayDomainFunctionResponseBody struct {
	DomainConfigList []*GetGrayDomainFunctionResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// example.com
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Progression *string `json:"Progression,omitempty" xml:"Progression,omitempty"`
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGrayDomainFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGrayDomainFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrayDomainFunctionResponseBody) GetDomainConfigList() []*GetGrayDomainFunctionResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *GetGrayDomainFunctionResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *GetGrayDomainFunctionResponseBody) GetProgression() *string {
	return s.Progression
}

func (s *GetGrayDomainFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGrayDomainFunctionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetGrayDomainFunctionResponseBody) SetDomainConfigList(v []*GetGrayDomainFunctionResponseBodyDomainConfigList) *GetGrayDomainFunctionResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *GetGrayDomainFunctionResponseBody) SetDomainName(v string) *GetGrayDomainFunctionResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBody) SetProgression(v string) *GetGrayDomainFunctionResponseBody {
	s.Progression = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBody) SetRequestId(v string) *GetGrayDomainFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBody) SetStatus(v string) *GetGrayDomainFunctionResponseBody {
	s.Status = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBody) Validate() error {
	if s.DomainConfigList != nil {
		for _, item := range s.DomainConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGrayDomainFunctionResponseBodyDomainConfigList struct {
	// example:
	//
	// 6295
	ConfigId     *int64                                                           `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs []*GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// example:
	//
	// ali_auth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// 222728944812032
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGrayDomainFunctionResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetGrayDomainFunctionResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) GetFunctionArgs() []*GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs {
	return s.FunctionArgs
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) GetParentId() *string {
	return s.ParentId
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) GetStatus() *string {
	return s.Status
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) SetConfigId(v int64) *GetGrayDomainFunctionResponseBodyDomainConfigList {
	s.ConfigId = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) SetFunctionArgs(v []*GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) *GetGrayDomainFunctionResponseBodyDomainConfigList {
	s.FunctionArgs = v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) SetFunctionName(v string) *GetGrayDomainFunctionResponseBodyDomainConfigList {
	s.FunctionName = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) SetParentId(v string) *GetGrayDomainFunctionResponseBodyDomainConfigList {
	s.ParentId = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) SetStatus(v string) *GetGrayDomainFunctionResponseBodyDomainConfigList {
	s.Status = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigList) Validate() error {
	if s.FunctionArgs != nil {
		for _, item := range s.FunctionArgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs struct {
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// example:
	//
	// req
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) GoString() string {
	return s.String()
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) SetArgName(v string) *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) SetArgValue(v string) *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *GetGrayDomainFunctionResponseBodyDomainConfigListFunctionArgs) Validate() error {
	return dara.Validate(s)
}
